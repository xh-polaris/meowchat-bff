package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewNewsLogic {
	return &NewNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewNewsLogic) NewNews(req *types.NewNewsReq) (resp *types.NewNewsResp, err error) {
	resp = new(types.NewNewsResp)

	if req.Id == "" {
		if err = checkCommunityPermission(l.ctx, l.svcCtx, req.CommunityId); err != nil {
			return
		}
		data, err := l.svcCtx.SystemRPC.CreateNews(l.ctx, &pb.CreateNewsReq{
			CommunityId: req.CommunityId,
			ImageUrl:    req.ImageUrl,
			LinkUrl:     req.LinkUrl,
			Type:        req.Type,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		if err = checkNewsPermission(l.ctx, l.svcCtx, req.Id); err != nil {
			return
		}
		_, err = l.svcCtx.SystemRPC.UpdateNews(l.ctx, &pb.UpdateNewsReq{
			Id:       req.Id,
			ImageUrl: req.ImageUrl,
			LinkUrl:  req.LinkUrl,
			Type:     req.Type,
		})
		if err != nil {
			return nil, err
		}
	}
	return
}
