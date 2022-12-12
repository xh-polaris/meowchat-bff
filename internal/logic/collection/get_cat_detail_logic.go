package collection

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCatDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCatDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatDetailLogic {
	return &GetCatDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCatDetailLogic) GetCatDetail(req *types.GetCatDetailReq) (resp *types.GetCatDetailResp, err error) {
	resp = new(types.GetCatDetailResp)
	data, err := l.svcCtx.CollectionRPC.RetrieveCat(l.ctx, &pb.RetrieveCatReq{
		CatId: req.CatId,
	})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Cat, data.Cat)
	if err != nil {
		return nil, err
	}

	return
}
