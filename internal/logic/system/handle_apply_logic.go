package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleApplyLogic {
	return &HandleApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleApplyLogic) HandleApply(req *types.HandleApplyReq) (resp *types.HandleApplyResp, err error) {
	resp = new(types.HandleApplyResp)
	_, err = l.svcCtx.SystemRPC.HandleApply(l.ctx, &pb.HandleApplyReq{
		ApplyId:    req.ApplyId,
		IsRejected: req.IsRejected,
	})
	if err != nil {
		return nil, err
	}
	return
}
