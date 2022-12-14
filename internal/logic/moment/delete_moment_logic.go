package moment

import (
	"context"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

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
	resp = new(types.DeleteMomentResp)
	_, err = l.svcCtx.MomentRPC.DeleteMoment(l.ctx, &pb.DeleteMomentReq{MomentId: req.MomentId})
	if err != nil {
		return nil, err
	}

	return
}
