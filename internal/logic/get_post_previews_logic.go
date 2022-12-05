package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostPreviewsLogic {
	return &GetPostPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostPreviewsLogic) GetPostPreviews(req *types.GetPostPreviewsReq) (resp *types.GetPostPreviewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
