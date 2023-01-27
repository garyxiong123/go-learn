package handler

import (
	"fmt"
	"net/http"

	"go-zero-demo/greet/internal/logic"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
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
		resp, err := l.Greet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
