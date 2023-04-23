package post

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *GetPostPreviewsLogic) makeRequest(req *types.GetPostPreviewsReq) *pb.ListPostReq {
	r := &pb.ListPostReq{}

	if req.SearchOptions != nil {
		if req.SearchOptions.Key != nil {
			r.SearchOptions = &pb.SearchOptions{
				Query: &pb.SearchOptions_AllFieldsKey{
					AllFieldsKey: *req.SearchOptions.Key,
				},
			}
		} else {
			r.SearchOptions = &pb.SearchOptions{
				Query: &pb.SearchOptions_MultiFieldsKey{
					MultiFieldsKey: &pb.SearchField{
						Text:  req.SearchOptions.Text,
						Title: req.SearchOptions.Title,
						Tag:   req.SearchOptions.Tag,
					},
				},
			}
		}
	}
	r.PaginationOptions = &pb.PaginationOptions{
		Offset:    req.Offset,
		Limit:     req.Limit,
		Backward:  req.Backward,
		LastToken: req.LastToken,
	}
	r.FilterOptions = &pb.FilterOptions{
		OnlyOfficial: req.OnlyOfficial,
		OnlyUserId:   req.OnlyUserId,
	}

	return r
}

func (l *GetPostPreviewsLogic) GetPostPreviews(req *types.GetPostPreviewsReq) (resp *types.GetPostPreviewsResp, err error) {
	resp = new(types.GetPostPreviewsResp)

	data, err := l.svcCtx.PostRPC.ListPost(l.ctx, l.makeRequest(req))
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Posts = make([]types.Post, len(data.Posts))
	for i, val := range data.Posts {
		respPost, _ := toRespPost(l.ctx, l.svcCtx, val)
		resp.Posts[i] = respPost
	}
	resp.Token = data.Token

	return
}
