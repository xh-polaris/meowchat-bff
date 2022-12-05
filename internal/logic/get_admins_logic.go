package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminsLogic {
	return &GetAdminsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminsLogic) GetAdmins() (resp *types.GetAdminsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
