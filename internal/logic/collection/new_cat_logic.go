package collection

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewCatLogic {
	return &NewCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewCatLogic) NewCat(req *types.NewCatReq) (resp *types.NewCatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
