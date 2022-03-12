package grpchandler

import (
	"github.com/Javellin01/go_course/internal/domain/usecase"
	"github.com/Javellin01/go_course/internal/presenters/grpc/pb"
)

type PlatformsServer struct {
	pb.PlatformsServiceServer
	usecase usecase.Usecase
}

func NewPlatformsServer(uc usecase.Usecase) PlatformsServer {
	return PlatformsServer{usecase: uc}
}
