package main

import (
	"flag"
	"fmt"
	"github.com/garyxiong123/go-learn/web/go-zero/basic/internal/svc"

	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var configFile = flag.String("f", "/Users/user/sdk/go_path/src/go-learn/web/go-zero/go-zero-demo/greet/etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			println("gary execute before -----------")
			next(writer, request)
			println("gary execute after -----------", []byte("sss"))

		}
	})

	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("swagger-ui-dist"))))

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	server.Start()

}
