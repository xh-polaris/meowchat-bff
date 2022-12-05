package comment

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
