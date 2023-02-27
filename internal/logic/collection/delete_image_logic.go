package collection

import (
	"context"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteImageLogic {
	return &DeleteImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteImageLogic) DeleteImage(req *types.DeleteImageReq) (resp *types.DeleteImageResp, err error) {
	data := pb.DeleteImageReq{ImageId: req.Id}
	_, err = l.svcCtx.CollectionRPC.DeleteImage(l.ctx, &data)
	if err != nil {
		return
	}
	resp = &types.DeleteImageResp{}
	return
}
