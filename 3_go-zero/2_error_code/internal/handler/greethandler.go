package handler

import (
	"fmt"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/logic"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/svc"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request

		defer func() {
			if x := recover(); x != nil {
				fmt.Printf("[ERROR]: My panic handle error: %s\n", x)
			}
		}()

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), ctx)
		resp, _ := l.Greet(req)
		httpx.OkJson(w, resp)
	}
}
