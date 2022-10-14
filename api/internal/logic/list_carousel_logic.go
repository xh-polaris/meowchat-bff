package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCarouselLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCarouselLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCarouselLogic {
	return &ListCarouselLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCarouselLogic) ListCarousel(req *types.ListCarouselReq) (resp *types.ListCarouselResp, err error) {
	// todo: add your logic here and delete this line

	return
}
