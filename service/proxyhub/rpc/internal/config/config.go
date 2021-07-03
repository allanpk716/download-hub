package config

import (
	"github.com/allanpk716/Downloadhub/service/proxyhub/model"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheRedis cache.ClusterConf	// 这里定义的是配置文件中需要读取的信息
	ProxyConf model.ProxyConf		// 自定义的代理配置
}
