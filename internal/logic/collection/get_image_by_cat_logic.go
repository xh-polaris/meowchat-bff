package collection

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImageByCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetImageByCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImageByCatLogic {
	return &GetImageByCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetImageByCatLogic) GetImageByCat(req *types.GetImageByCatReq) (resp *types.GetImageByCatResp, err error) {

	data := pb.GetImageByCatReq{
		CatId: req.CatId,
		Limit: req.Limit,
	}
	if req.PrevId != "" {
		data.PrevId = &req.PrevId
	}
	res, err := l.svcCtx.CollectionRPC.GetImageByCat(l.ctx, &data)
	if err != nil {
		return nil, err
	}
	resp = &types.GetImageByCatResp{
		LastId:   res.LastId,
		ImageUrl: res.ImageUrl,
	}
	return
}
