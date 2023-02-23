package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewCommunityLogic {
	return &NewCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewCommunityLogic) NewCommunity(req *types.NewCommunityReq) (resp *types.NewCommunityResp, err error) {
	resp = new(types.NewCommunityResp)

	if req.Id == "" {
		data, err := l.svcCtx.SystemRPC.CreateCommunity(l.ctx, &pb.CreateCommunityReq{
			Name:     req.Name,
			ParentId: req.ParentId,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		_, err = l.svcCtx.SystemRPC.UpdateCommunity(l.ctx, &pb.UpdateCommunityReq{
			Id:       req.Id,
			Name:     req.Name,
			ParentId: req.ParentId,
		})
		if err != nil {
			return nil, err
		}
	}
	return
}
