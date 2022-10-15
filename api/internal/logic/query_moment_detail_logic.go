package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMomentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMomentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMomentDetailLogic {
	return &QueryMomentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMomentDetailLogic) QueryMomentDetail(req *types.QueryMomentReq) (resp *types.QueryMomentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
