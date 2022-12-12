package collection

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCatLogic {
	return &DeleteCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCatLogic) DeleteCat(req *types.DeleteCatReq) (resp *types.DeleteCatResp, err error) {
	resp = new(types.DeleteCatResp)
	_, err = l.svcCtx.CollectionRPC.DeleteCat(l.ctx, &pb.DeleteCatReq{CatId: req.CatId})
	if err != nil {
		return nil, err
	}

	return
}
