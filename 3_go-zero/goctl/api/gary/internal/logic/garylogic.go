package logic

import (
	"context"

	"github.com/garyxiong123/go-learn/web/go-zero/goctl/api/gary/internal/svc"
	"github.com/garyxiong123/go-learn/web/go-zero/goctl/api/gary/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GaryLogic {
	return &GaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GaryLogic) Gary(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
