package collection

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCatPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCatPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatPreviewsLogic {
	return &GetCatPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCatPreviewsLogic) GetCatPreviews() (resp *types.GetCatPreviewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
