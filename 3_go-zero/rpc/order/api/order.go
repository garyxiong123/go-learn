package main

import (
	"flag"
	"github.com/garyxiong123/go-learn/web/go-zero/rpc/order/api/internal/config"
	"github.com/garyxiong123/go-learn/web/go-zero/rpc/order/api/internal/handler"
	"github.com/garyxiong123/go-learn/web/go-zero/rpc/order/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	//fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
