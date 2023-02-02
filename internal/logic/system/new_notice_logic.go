package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewNoticeLogic {
	return &NewNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewNoticeLogic) NewNotice(req *types.NewNoticeReq) (resp *types.NewNoticeResp, err error) {
	resp = new(types.NewNoticeResp)

	if req.Id == "" {
		if err = checkCommunityPermission(l.ctx, l.svcCtx, req.CommunityId); err != nil {
			return
		}
		data, err := l.svcCtx.SystemRPC.CreateNotice(l.ctx, &pb.CreateNoticeReq{
			Text:        req.Text,
			CommunityId: req.CommunityId,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.Id
		if err = checkNoticePermission(l.ctx, l.svcCtx, req.Id); err != nil {
			return
		}
		_, err := l.svcCtx.SystemRPC.UpdateNotice(l.ctx, &pb.UpdateNoticeReq{
			Id:   req.Id,
			Text: req.Text,
		})
		if err != nil {
			return nil, err
		}
	}

	return
}
