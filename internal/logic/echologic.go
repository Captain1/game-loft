package logic

import (
	"context"

	"gameloft/internal/svc"
	"gameloft/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EchoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEchoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EchoLogic {
	return &EchoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EchoLogic) Echo(req *types.Request) (resp *types.Response, err error) {
	return resp, nil
}
