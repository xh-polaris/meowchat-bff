package post

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOwnPostPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOwnPostPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOwnPostPreviewsLogic {
	return &GetOwnPostPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOwnPostPreviewsLogic) GetOwnPostPreviews(req *types.GetOwnPostPreviewsReq) (resp *types.GetOwnPostPreviewsResp, err error) {
	resp = new(types.GetOwnPostPreviewsResp)

	data, err := l.svcCtx.PostRPC.ListPostByUserId(l.ctx, &pb.ListPostByUserIdReq{
		UserId: l.ctx.Value("userId").(string),
		Skip:   req.Page * pageSize,
		Count:  pageSize,
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
