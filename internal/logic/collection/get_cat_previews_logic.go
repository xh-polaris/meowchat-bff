package collection

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var pageSize int64 = 20

type GetCatPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCatPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatPreviewsLogic {
	return &GetCatPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCatPreviewsLogic) GetCatPreviews(req *types.GetCatPreviewsReq) (resp *types.GetCatPreviewsResp, err error) {
	resp = new(types.GetCatPreviewsResp)
	data, err := l.svcCtx.CollectionRPC.ListCat(l.ctx, &pb.ListCatReq{
		CommunityId: req.CommunityId,
		Count:       pageSize,
		Skip:        req.Page * pageSize,
	})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(resp.Cats, data.Cats)
	if err != nil {
		return nil, err
	}

	return
}
