package post

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewPostLogic {
	return &NewPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewPostLogic) NewPost(req *types.NewPostReq) (resp *types.NewPostResp, err error) {
	// todo: add your logic here and delete this line

	return
}
