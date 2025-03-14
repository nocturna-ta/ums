package controller

import (
	"github.com/nocturna-ta/api-gateway-grpc-lib/proto"
	"github.com/nocturna-ta/golib/grpc"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type GRPC struct {
	port   uint
	authUc usecases.AuthUseCases
}

type server struct {
	proto.AuthServiceServer
	authUc usecases.AuthUseCases
}

type Options struct {
	Port   uint
	AuthUc usecases.AuthUseCases
}

func New(opts *Options) *GRPC {
	return &GRPC{
		port:   opts.Port,
		authUc: opts.AuthUc,
	}
}

func (g *GRPC) Register() *grpc.Server {
	srv := grpc.NewServer(&grpc.ServerOptions{
		Port: g.port,
	})

	grpcServiceServer := &server{
		authUc: g.authUc,
	}

	srv.Register(proto.RegisterAuthServiceServer, grpcServiceServer)

	return srv
}
