package user

import (
	"context"
	"github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(*types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	resp = new(types.GetUserInfoResp)
	userId := l.ctx.Value("userId").(string)
	data, err := l.svcCtx.UserRPC.GetUser(l.ctx, &pb.GetUserReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	resp.User = types.UserPreview{
		Id:        data.UserId,
		Nickname:  data.Nickname,
		AvatarUrl: data.AvatarUrl,
	}
	return
}
