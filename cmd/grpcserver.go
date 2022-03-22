package main

import (
	"context"
	mongorepository "github.com/Javellin01/go_course/internal/data/repository/mongo"
	"github.com/Javellin01/go_course/internal/domain/usecase"
	grpchandler "github.com/Javellin01/go_course/internal/presenters/grpc"
	"github.com/Javellin01/go_course/internal/presenters/grpc/pb"
	"github.com/Javellin01/go_course/pkg/env"
	"github.com/Javellin01/go_course/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Config struct {
	Debug    bool   `env:"DEBUG,default=true"`
	Port     string `env:"LISTEN_PORT,default=8080"`
	MongoUri string `env:"MONGO_URI,default=mongodb://root:mongodb@localhost:27017"`
}

func RunGRPC() error {
	var config Config

	if err := env.Unmarshal(&config); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", net.JoinHostPort("", config.Port))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := connectDatabase(ctx, config.MongoUri)
	if err != nil {
		return err
	}

	campaignRepository := mongorepository.NewCampaignRepository(db, time.Second*3)
	advertiserRepository := mongorepository.NewAdvertiserRepository(db, time.Second*3)
	usecases := usecase.Usecase{
		Campaign:   usecase.NewCampaignUsecase(ctx, campaignRepository),
		Advertiser: usecase.NewAdvertiserUsecase(ctx, advertiserRepository),
	}

	handler := grpchandler.NewPlatformsServer(usecases)
	server := grpc.NewServer()

	pb.RegisterPlatformsServiceServer(server, handler)

	return server.Serve(listener)
}

func connectDatabase(ctx context.Context, uri string) (*mongo.Database, error) {
	mongoDB, err := mongodb.New(ctx, uri)
	db := mongoDB.Database("mongodb")
	if err != nil {
		return db, err
	}

	err = mongoDB.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
