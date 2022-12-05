package logic

import (
	"context"

	"github.com/xh-polaris/account-rpc/pb"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerifyCodeLogic {
	return &SendVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerifyCodeLogic) SendVerifyCode(req *types.SendVerifyCodeReq) (*types.SendVerifyCodeResp, error) {
	_, err := l.svcCtx.AccountRPC.SendVerifyCode(l.ctx, &pb.SendVerifyCodeReq{
		AuthType: req.AuthType,
		AuthId:   req.AuthId,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
