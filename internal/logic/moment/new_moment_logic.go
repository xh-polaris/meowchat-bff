package moment

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewMomentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewMomentLogic {
	return &NewMomentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewMomentLogic) NewMoment(req *types.NewMomentReq) (resp *types.NewMomentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
