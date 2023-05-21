package comment

import (
	"context"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	resp = new(types.DeleteCommentResp)
	_, err = l.svcCtx.CommentRPC.DeleteComment(l.ctx, &pb.DeleteCommentByIdRequest{Id: req.CommentId})
	if err != nil {
		return nil, err
	}

	return
}
