package logic

import (
	"context"
	"github.com/garyxiong123/go-learn/web/go-zero/common/errorx"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/svc"
	"github.com/garyxiong123/go-learn/web/go-zero/error_code/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
	//time.Sleep(5 * time.Second)
	if "errorcode" == req.Name {
		return nil, errorx.NewDefaultError("参数错误")
	}

	if "panic" == req.Name {
		panic("gary panic")
	}

	return &types.Response{
		Message: "Hello go-zero",
	}, nil
}
