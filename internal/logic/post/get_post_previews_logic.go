package post

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const pageSize = 10

type GetPostPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostPreviewsLogic {
	return &GetPostPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostPreviewsLogic) GetPostPreviews(req *types.GetPostPreviewsReq) (resp *types.GetPostPreviewsResp, err error) {
	resp = new(types.GetPostPreviewsResp)

	data, err := l.svcCtx.PostRPC.ListPost(l.ctx, &pb.ListPostReq{
		Skip:  req.Page * pageSize,
		Count: pageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Posts = make([]types.Post, len(data.Posts))
	for i, val := range data.Posts {
		respPost, _ := toRespPost(l.ctx, l.svcCtx, val)
		resp.Posts[i] = respPost
	}

	return
}
