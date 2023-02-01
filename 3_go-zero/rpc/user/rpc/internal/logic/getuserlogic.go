package logic

import (
	"context"
	"github.com/garyxiong123/go-learn/web/go-zero/rpc/user/rpc/internal/svc"
	"github.com/garyxiong123/go-learn/web/go-zero/rpc/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	return &user.UserResponse{
		Id:   "1",
		Name: "test",
	}, nil
}
