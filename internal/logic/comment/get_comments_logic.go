package comment

import (
	"context"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	// todo: add your logic here and delete this line

	return
}
