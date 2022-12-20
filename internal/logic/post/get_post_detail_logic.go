package post

import (
	"context"
	"github.com/jinzhu/copier"
	commentpb "github.com/xh-polaris/meowchat-comment-rpc/pb"
	"github.com/xh-polaris/meowchat-like-rpc/like"
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
	err = copier.Copy(&resp, post)
	if err != nil {
		return types.Post{}, err
	}

	// Tag
	resp.Tags = post.Tags

	// user preview
	user, err := svcCtx.UserRPC.GetUser(ctx, &userpb.GetUserReq{UserId: post.UserId})
	if user != nil && err == nil {
		resp.User = types.UserPreview{
			Id:        post.UserId,
			Nickname:  user.Nickname,
			AvatarUrl: user.AvatarUrl,
		}
	}

	// likes
	likes, err := svcCtx.LikeRPC.GetTargetLikes(ctx, &likepb.GetTargetLikesReq{
		TargetId: post.Id,
		Type:     like.TargetTypePost,
	})
	if likes != nil && err == nil {
		resp.Likes = likes.Count
	}

	// comments
	// TODO count
	comments, err := svcCtx.CommentRPC.ListCommentByParent(ctx, &commentpb.ListCommentByParentRequest{
		Type:     "post",
		ParentId: post.Id,
		Skip:     0,
		Limit:    9999,
	})
	if comments != nil && err == nil {
		resp.Comments = int64(len(comments.Comments))
	}

	return
}

func (l *GetPostDetailLogic) GetPostDetail(req *types.GetPostDetailReq) (resp *types.GetPostDetailResp, err error) {
	resp = new(types.GetPostDetailResp)

	data, err := l.svcCtx.PostRPC.RetrievePost(l.ctx, &pb.RetrievePostReq{PostId: req.PostId})
	if err != nil {
		return nil, err
	}

	respPost, _ := toRespPost(l.ctx, l.svcCtx, data.GetPost())
	resp.Post = respPost

	return
}
