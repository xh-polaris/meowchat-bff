package user

import (
	"context"
	"github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, err error) {
	resp = new(types.UpdateUserInfoResp)
	userId := l.ctx.Value("userId").(string)
	_, err = l.svcCtx.UserRPC.UpdateUser(l.ctx, &pb.UpdateUserReq{
		UserId:    userId,
		AvatarUrl: req.AvatarUrl,
		Nickname:  req.Nickname,
	})
	if err != nil {
		return nil, err
	}

	return
}
