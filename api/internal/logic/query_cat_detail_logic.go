package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCatDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCatDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCatDetailLogic {
	return &QueryCatDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCatDetailLogic) QueryCatDetail(req *types.QueryCatReq) (resp *types.QueryCatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
