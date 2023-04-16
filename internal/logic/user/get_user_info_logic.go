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

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	resp = new(types.GetUserInfoResp)

	var userId string
	if req.UserId != nil {
		userId = *req.UserId
	} else {
		userId = l.ctx.Value("userId").(string)
	}

	data, err := l.svcCtx.UserRPC.GetUserDetail(l.ctx, &pb.GetUserDetailReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	resp.User = types.User{
		Id:        data.User.Id,
		Nickname:  data.User.Nickname,
		AvatarUrl: data.User.AvatarUrl,
		Motto:     data.User.Motto,
	}
	return
}
