package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminDetailLogic {
	return &GetAdminDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminDetailLogic) GetAdminDetail(req *types.GetAdminDetailReq) (resp *types.GetAdminDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
