package handler

import (
	"net/http"

	"github.com/garyxiong123/go-learn/web/go-zero/rate_limit/internal/logic"
	"github.com/garyxiong123/go-learn/web/go-zero/rate_limit/internal/svc"
	"github.com/garyxiong123/go-learn/web/go-zero/rate_limit/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), svcCtx)
		resp, err := l.Greet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
