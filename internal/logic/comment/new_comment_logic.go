package comment

import (
	"context"
	"github.com/xh-polaris/meowchat-bff/internal/logic/util"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	WechatSecUrl = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
)

type NewCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewCommentLogic {
	return &NewCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewCommentLogic) NewComment(req *types.NewCommentReq) (resp *types.NewCommentResp, err error) {
	resp = new(types.NewCommentResp)
	userId := l.ctx.Value("userId").(string)
	openId := l.ctx.Value("openId").(string)

	err = util.MsgSecCheck(l.ctx, l.svcCtx, req.Text, openId, 2)
	if err != nil {
		return nil, err
	}
	// 获取回复用户id
	replyToId := ""
	if req.Scope == "comment" {
		replyTo, err := l.svcCtx.CommentRPC.RetrieveCommentById(l.ctx, &pb.RetrieveCommentByIdRequest{Id: req.Id})
		if err != nil {
			return nil, err
		}
		replyToId = replyTo.Comment.AuthorId
	}

	_, err = l.svcCtx.CommentRPC.CreateComment(l.ctx, &pb.CreateCommentRequest{
		Text:     req.Text,
		AuthorId: userId,
		ReplyTo:  replyToId,
		Type:     req.Scope,
		ParentId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return
}
