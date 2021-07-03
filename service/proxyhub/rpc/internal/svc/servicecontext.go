package svc

import (
	config2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/internal/config"
)

type ServiceContext struct {
	Config config2.Config
}

func NewServiceContext(c config2.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
