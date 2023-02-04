package system

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommunityLogic {
	return &ListCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommunityLogic) ListCommunity(req *types.ListCommunityReq) (resp *types.ListCommunityResp, err error) {
	resp = new(types.ListCommunityResp)

	data, err := l.svcCtx.SystemRPC.ListCommunity(l.ctx, &pb.ListCommunityReq{ParentId: req.ParentId, Size: -1})
	if err != nil {
		return nil, err
	}
	resp.Communities = make([]types.Community, len(data.Communities))

	err = copier.Copy(&resp.Communities, &data.Communities)
	if err != nil {
		return nil, err
	}

	return
}
