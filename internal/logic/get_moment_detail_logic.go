package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-api/internal/svc"
	"github.com/xh-polaris/cat-community-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMomentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMomentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMomentDetailLogic {
	return &GetMomentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMomentDetailLogic) GetMomentDetail(req *types.GetMomentDetailReq) (resp *types.GetMomentDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
