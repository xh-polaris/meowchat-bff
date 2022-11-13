package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-api/internal/svc"
	"github.com/xh-polaris/cat-community-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMomentPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMomentPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMomentPreviewsLogic {
	return &GetMomentPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMomentPreviewsLogic) GetMomentPreviews(req *types.GetMomentPreviewsReq) (resp *types.GetMomentPreviewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}