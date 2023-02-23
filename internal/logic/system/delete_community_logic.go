package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommunityLogic {
	return &DeleteCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommunityLogic) DeleteCommunity(req *types.DeleteCommunityReq) (resp *types.DeleteCommunityResp, err error) {
	resp = new(types.DeleteCommunityResp)

	_, err = l.svcCtx.SystemRPC.DeleteCommunity(l.ctx, &pb.DeleteCommunityReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return
}
