package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewsLogic {
	return &GetNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNewsLogic) GetNews() (resp *types.GetNewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
