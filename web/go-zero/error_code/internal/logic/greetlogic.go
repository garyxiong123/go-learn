package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-demo/greet/common/errorx"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
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
	if "code" == req.Name {
		return nil, errorx.NewDefaultError("参数错误")
	}

	if "panic" == req.Name {
		panic("gary panic")
	}

	return &types.Response{
		Message: "Hello go-zero",
	}, nil
}
