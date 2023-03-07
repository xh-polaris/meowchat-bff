package comment

import (
	"context"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"
	userpb "github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const pageSize = 10

type GetCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsLogic) GetComments(req *types.GetCommentsReq) (resp *types.GetCommentsResp, err error) {
	resp = new(types.GetCommentsResp)

	data, err := l.svcCtx.CommentRPC.ListCommentByParent(l.ctx, &pb.ListCommentByParentRequest{
		ParentId: req.Id,
		Type:     req.Scope,
		Skip:     req.Page * pageSize,
		Limit:    pageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Comments = make([]types.Comment, len(data.Comments))
	for i, comment := range data.Comments {
		// 评论作者信息
		var author types.UserPreview
		author.Id = comment.AuthorId
		// 暂时不处理error
		user, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userpb.GetUserReq{
			UserId: comment.AuthorId,
		})
		if err != nil {
			return nil, err
		}
		if user != nil && err == nil {
			author.Nickname = user.User.Nickname
			author.AvatarUrl = user.User.AvatarUrl
		}

		// 回复对象用户名
		replyName := ""
		if comment.ReplyTo != "" {
			replyToUser, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userpb.GetUserReq{
				UserId: comment.ReplyTo,
			})
			if replyToUser != nil && err == nil {
				replyName = replyToUser.User.Nickname
			}
		}

		// 子评论数量
		// TODO count rpc
		count := 0
		children, err := l.svcCtx.CommentRPC.ListCommentByParent(l.ctx, &pb.ListCommentByParentRequest{
			Type:     "comment",
			ParentId: comment.Id,
			Skip:     0,
			Limit:    9999,
		})
		if children != nil && err == nil {
			count = len(children.Comments)
		}

		resp.Comments[i] = types.Comment{
			Id:        comment.Id,
			CreateAt:  comment.CreateAt,
			Text:      comment.Text,
			User:      author,
			Comments:  int64(count),
			ReplyName: replyName,
		}
	}

	return
}
