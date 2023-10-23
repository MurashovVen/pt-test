package controller

import (
	"pt-test/pkg/app/configuration"
	"pt-test/pkg/app/logger"
)

type Option func(srv *Server)

func WithAddress(addr string) Option {
	return func(srv *Server) {
		srv.server.Addr = addr
	}
}

func WithEnvironment(env configuration.Environment) Option {
	return func(srv *Server) {
		srv.env = env
	}
}

func WithLogger(log *logger.Logger) Option {
	return func(srv *Server) {
		srv.logger = log.Named(srv.Name())
	}
}
