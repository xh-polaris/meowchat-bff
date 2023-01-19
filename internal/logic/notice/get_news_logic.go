package notice

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewsLogic {
	return &GetNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNewsLogic) GetNews(req *types.GetNewsReq) (resp *types.GetNewsResp, err error) {
	resp = new(types.GetNewsResp)
	resp.News = make([]types.News, 0)

	data, err := l.svcCtx.SystemRPC.ListNews(l.ctx, &pb.ListNewsReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.News, &data.News)
	if err != nil {
		return nil, err
	}

	return
}
