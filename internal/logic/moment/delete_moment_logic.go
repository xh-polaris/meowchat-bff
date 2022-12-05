package moment

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMomentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMomentLogic {
	return &DeleteMomentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMomentLogic) DeleteMoment(req *types.DeleteMomentReq) (resp *types.DeleteMomentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
