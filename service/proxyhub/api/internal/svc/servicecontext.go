package svc

import (
	config2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/config"
	getter2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/getter"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config2.Config
	Getter getter2.Getter
}

func NewServiceContext(c config2.Config) *ServiceContext {
	//缓存
	return &ServiceContext{
		Config: c,
		Getter: getter2.NewGetter(zrpc.MustNewClient(c.Get)),
	}
}
