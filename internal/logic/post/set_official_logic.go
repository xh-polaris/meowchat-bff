package post

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetOfficialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetOfficialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetOfficialLogic {
	return &SetOfficialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetOfficialLogic) SetOfficial(req *types.SetOfficialReq) (resp *types.SetOfficialResp, err error) {
	resp = new(types.SetOfficialResp)
	_, err = l.svcCtx.PostRPC.SetOfficial(l.ctx, &pb.SetOfficialReq{
		PostId:   req.PostId,
		IsRemove: req.IsRemove,
	})
	if err != nil {
		return nil, err
	}

	return
}
