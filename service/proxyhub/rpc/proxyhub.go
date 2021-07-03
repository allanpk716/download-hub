package main

import (
	"flag"
	"fmt"
	ProxyHub2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/ProxyHub"
	config2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/internal/config"
	server2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/internal/server"
	svc2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/proxyhub.yaml", "the config file")

func main() {
	flag.Parse()

	var c config2.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc2.NewServiceContext(c)
	srv := server2.NewGetterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ProxyHub2.RegisterGetterServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
