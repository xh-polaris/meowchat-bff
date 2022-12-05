package collection

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCatDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCatDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatDetailLogic {
	return &GetCatDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCatDetailLogic) GetCatDetail(req *types.GetCatDetailReq) (resp *types.GetCatDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
