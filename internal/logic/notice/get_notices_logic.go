package notice

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoticesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticesLogic {
	return &GetNoticesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticesLogic) GetNotices(req *types.GetNoticesReq) (resp *types.GetNoticesResp, err error) {
	resp = new(types.GetNoticesResp)
	resp.Notices = make([]types.Notice, 0)

	data, err := l.svcCtx.SystemRPC.ListNotice(l.ctx, &pb.ListNoticeReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Notices, &data.Notices)
	if err != nil {
		return nil, err
	}

	return
}
