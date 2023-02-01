package handler

import (
	"net/http"

	"github.com/garyxiong123/go-learn/web/go-zero/goctl/api/gary/internal/logic"
	"github.com/garyxiong123/go-learn/web/go-zero/goctl/api/gary/internal/svc"
	"github.com/garyxiong123/go-learn/web/go-zero/goctl/api/gary/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGaryLogic(r.Context(), svcCtx)
		resp, err := l.Gary(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
