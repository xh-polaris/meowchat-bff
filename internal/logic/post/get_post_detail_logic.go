package post

import (
	"context"
	commentpb "github.com/xh-polaris/meowchat-comment-rpc/pb"
	"github.com/xh-polaris/meowchat-like-rpc/likerpc"
	likepb "github.com/xh-polaris/meowchat-like-rpc/pb"
	"github.com/xh-polaris/meowchat-post-rpc/pb"
	userpb "github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func toRespPost(ctx context.Context, svcCtx *svc.ServiceContext, post *pb.Post) (resp types.Post, err error) {
	resp = types.Post{
		Id:         post.Id,
		CreateAt:   post.CreateAt,
		Title:      post.Title,
		Text:       post.Text,
		CoverUrl:   post.CoverUrl,
		Tags:       post.Tags,
		IsOfficial: post.IsOfficial,
	}

	// user preview
	user, err := svcCtx.UserRPC.GetUser(ctx, &userpb.GetUserReq{UserId: post.UserId})
	if user != nil && err == nil {
		resp.User = types.UserPreview{
			Id:        post.UserId,
			Nickname:  user.User.Nickname,
			AvatarUrl: user.User.AvatarUrl,
		}
	}

	// likes
	likes, err := svcCtx.LikeRPC.GetTargetLikes(ctx, &likepb.GetTargetLikesReq{
		TargetId: post.Id,
		Type:     likerpc.TargetTypePost,
	})
	if likes != nil && err == nil {
		resp.Likes = likes.Count
	}

	// comments
	data, err := svcCtx.CommentRPC.CountCommentByParent(ctx, &commentpb.CountCommentByParentRequest{
		Type:     "post",
		ParentId: post.Id,
	})
	if err == nil {
		resp.Comments = data.Total
	}

	return
}

func (l *GetPostDetailLogic) GetPostDetail(req *types.GetPostDetailReq) (resp *types.GetPostDetailResp, err error) {
	resp = new(types.GetPostDetailResp)

	data, err := l.svcCtx.PostRPC.RetrievePost(l.ctx, &pb.RetrievePostReq{PostId: req.PostId})
	if err != nil {
		return nil, err
	}

	respPost, _ := toRespPost(l.ctx, l.svcCtx, data.Post)
	resp.Post = respPost

	return
}
