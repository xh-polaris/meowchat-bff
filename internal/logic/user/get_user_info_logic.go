package user

import (
	"context"
	pb2 "github.com/xh-polaris/meowchat-moment-rpc/pb"
	pb3 "github.com/xh-polaris/meowchat-post-rpc/pb"
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

func (l *GetUserInfoLogic) getLessDependentInfo(user *types.User) {
	momentCount, err := l.svcCtx.MomentRPC.CountMoment(l.ctx, &pb2.CountMomentReq{
		FilterOptions: &pb2.FilterOptions{OnlyUserId: &user.Id},
	})
	if err != nil {
		logx.Error(err)
	}
	user.Article += momentCount.Total

	postCount, err := l.svcCtx.PostRPC.CountPost(l.ctx, &pb3.CountPostReq{
		FilterOptions: &pb3.FilterOptions{
			OnlyUserId: &user.Id,
		},
	})
	if err != nil {
		logx.Error(err)
	}
	user.Article += postCount.Total
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

	l.getLessDependentInfo(&resp.User)

	return
}
