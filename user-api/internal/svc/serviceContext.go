package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	TetsMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TetsMiddleware: middleware.NewTetsMiddleware().Handle,
	}
}
