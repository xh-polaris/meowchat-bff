package notice

import (
	"context"
	"github.com/xh-polaris/meowchat-notice-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNewsLogic {
	return &DeleteNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNewsLogic) DeleteNews(req *types.DeleteNewsReq) (resp *types.DeleteNewsResp, err error) {
	resp = new(types.DeleteNewsResp)

	_, err = l.svcCtx.NoticeRPC.DeleteNews(l.ctx, &pb.DeleteNewsReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return
}
