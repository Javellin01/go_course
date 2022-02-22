package main

import (
	"github.com/Javellin01/go_course/pkg/env"
	"google.golang.org/grpc"
	"net"
)

type Config struct {
	Debug bool   `env:"DEBUG,default=true"`
	Port  string `env:"LISTEN_PORT,default=8080"`
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

	server := grpc.NewServer()

	return server.Serve(listener)
}
