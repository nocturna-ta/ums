package grpc

import (
	"github.com/nocturna-ta/golib/grpc"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/handler/grpc/controller"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Options struct {
	Cfg    config.MainConfig
	AuthUc usecases.AuthUseCases
}

type Handler struct {
	opts       *Options
	grpcServer *grpc.Server
}

func New(opts *Options) *Handler {
	handler := &Handler{
		opts: opts,
	}

	handler.grpcServer = controller.New(&controller.Options{
		Port:   opts.Cfg.GrpcServer.Port,
		AuthUc: opts.AuthUc,
	}).Register()

	return handler
}
func (h *Handler) Run() {
	h.grpcServer.MustStart()
}

func (h *Handler) Stop() {
	h.grpcServer.Stop()
}
