package sts

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/sts-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyTokenLogic {
	return &ApplyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyTokenLogic) ApplyToken(req *types.ApplyTokenReq) (resp *types.ApplyTokenResp, err error) {
	resp = new(types.ApplyTokenResp)
	userId := l.ctx.Value("userId").(string)
	data, err := l.svcCtx.StsRPC.GetUserCosSts(l.ctx, &pb.GetUserCosStsReq{
		UserId: userId,
		Path:   req.Path,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(resp, data)
	if err != nil {
		return nil, err
	}
	return
}
