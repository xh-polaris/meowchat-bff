package logic

import (
	"context"

	"github.com/xh-polaris/account-rpc/pb"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPasswordLogic {
	return &SetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetPasswordLogic) SetPassword(req *types.SetPasswordReq) (*types.SetPasswordResp, error) {
	userId := l.ctx.Value("userId").(string)
	_, err := l.svcCtx.AccountRPC.SetPassword(l.ctx, &pb.SetPasswordReq{
		UserId:   userId,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
