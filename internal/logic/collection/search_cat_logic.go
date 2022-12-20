package collection

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCatLogic {
	return &SearchCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCatLogic) SearchCat(req *types.SearchCatReq) (resp *types.SearchCatResp, err error) {
	resp = new(types.SearchCatResp)
	data, err := l.svcCtx.CollectionRPC.SearchCat(l.ctx, &pb.SearchCatReq{
		CommunityId: req.CommunityId,
		Count:       pageSize,
		Skip:        req.Page * pageSize,
		Keyword:     req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	resp.Cats = make([]types.CatPreview, 0, pageSize)
	err = copier.Copy(&resp.Cats, data.Cats)
	for i := 0; i < len(resp.Cats); i++ {
		if len(data.Cats[i].Avatars) > 0 {
			resp.Cats[i].AvatarUrl = data.Cats[i].Avatars[0]
		}
	}
	if err != nil {
		return nil, err
	}

	return
}
