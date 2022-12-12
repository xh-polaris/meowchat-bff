package collection

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewCatLogic {
	return &NewCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewCatLogic) NewCat(req *types.NewCatReq) (resp *types.NewCatResp, err error) {
	resp = new(types.NewCatResp)
	cat := new(pb.Cat)
	err = copier.Copy(cat, req)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		var data *pb.CreateCatResp
		data, err = l.svcCtx.CollectionRPC.CreateCat(l.ctx, &pb.CreateCatReq{Cat: cat})
		resp.CatId = data.CatId
	} else {
		_, err = l.svcCtx.CollectionRPC.UpdateCat(l.ctx, &pb.UpdateCatReq{Cat: cat})
		resp.CatId = cat.Id
	}

	if err != nil {
		return nil, err
	}

	return
}
