package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewAdminLogic {
	return &NewAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewAdminLogic) NewAdmin(req *types.NewAdminReq) (resp *types.NewAdminResp, err error) {
	resp = new(types.NewAdminResp)
	if req.Id == "" {
		data, err := l.svcCtx.SystemRPC.CreateAdmin(l.ctx, &pb.CreateAdminReq{
			CommunityId: req.CommunityId,
			Name:        req.Name,
			Title:       req.Title,
			Phone:       req.Phone,
			Wechat:      req.Wechat,
			AvatarUrl:   req.AvatarUrl,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		_, err := l.svcCtx.SystemRPC.UpdateAdmin(l.ctx, &pb.UpdateAdminReq{
			Id:          req.Id,
			CommunityId: req.CommunityId,
			Name:        req.Name,
			Title:       req.Title,
			Phone:       req.Phone,
			Wechat:      req.Wechat,
			AvatarUrl:   req.AvatarUrl,
		})
		if err != nil {
			return nil, err
		}
	}
	return
}
