// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"github.com/garyxiong123/go-learn/web/go-zero/basic/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: GreetHandler(serverCtx),
			},
		},
	)
}
