package svc

import (
	"github.com/garyxiong123/go-learn/web/go-zero/basic/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
