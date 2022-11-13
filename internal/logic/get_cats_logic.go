package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-api/internal/svc"
	"github.com/xh-polaris/cat-community-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatsLogic {
	return &GetCatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCatsLogic) GetCats() (resp *types.GetCatPreviewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
