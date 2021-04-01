package main

import (
	"flag"
	"fmt"

	"github.com/allanpk716/Downloadhub/service/onlineVideoDl/cmd/rpc/internal/config"
	"github.com/allanpk716/Downloadhub/service/onlineVideoDl/cmd/rpc/internal/server"
	"github.com/allanpk716/Downloadhub/service/onlineVideoDl/cmd/rpc/internal/svc"
	"github.com/allanpk716/Downloadhub/service/onlineVideoDl/cmd/rpc/onlineVideoDl"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/onlinevideodl.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOnlineVideoDlServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		onlineVideoDl.RegisterOnlineVideoDlServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
