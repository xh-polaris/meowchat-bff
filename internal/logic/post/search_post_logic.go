package post

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPostLogic {
	return &SearchPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchPostLogic) SearchPost(req *types.SearchPostReq) (resp *types.SearchPostResp, err error) {
	resp = new(types.SearchPostResp)

	data, err := l.svcCtx.PostRPC.SearchPost(l.ctx, &pb.SearchPostReq{
		Keyword: req.Keyword,
		Skip:    req.Page * pageSize,
		Count:   pageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Posts = make([]types.Post, len(data.GetPosts()))
	for i, val := range data.GetPosts() {
		respPost, _ := toRespPost(l.ctx, l.svcCtx, val)
		resp.Posts[i] = respPost
	}

	return
}
