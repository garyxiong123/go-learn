package main

import (
	"flag"
	"github.com/garyxiong123/go-learn/web/go-zero/common/errorx"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/config"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/handler"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"net/http"
)

var configFile = flag.String("f", "./etc/greet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	//fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {

		return func(writer http.ResponseWriter, request *http.Request) {
			println("gary execute before -----------")

			next(writer, request) //filter

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
