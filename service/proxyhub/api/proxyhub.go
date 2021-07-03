package main

import (
	"flag"
	"fmt"
	config2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/config"
	handler2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/handler"
	svc2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/ProxyHub-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config2.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc2.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler2.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
