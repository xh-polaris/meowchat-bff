package notice

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoticeLogic {
	return &DeleteNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNoticeLogic) DeleteNotice(req *types.DeleteNoticeReq) (resp *types.DeleteNoticeResp, err error) {
	resp = new(types.DeleteNoticeResp)

	if err = checkNoticePermission(l.ctx, l.svcCtx, req.Id); err != nil {
		return
	}

	_, err = l.svcCtx.SystemRPC.DeleteNotice(l.ctx, &pb.DeleteNoticeReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return
}
