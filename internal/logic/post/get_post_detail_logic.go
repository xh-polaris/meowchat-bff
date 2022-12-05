package post

import (
	"context"

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

func (l *GetPostDetailLogic) GetPostDetail(req *types.GetPostDetailReq) (resp *types.GetPostDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
