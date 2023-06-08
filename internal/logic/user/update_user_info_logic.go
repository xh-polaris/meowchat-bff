package user

import (
	"context"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/xh-polaris/meowchat-user-rpc/pb"
	"net/url"

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
	openId := l.ctx.Value("openId").(string)

	err = util.MsgSecCheck(l.ctx, l.svcCtx, req.Nickname, openId, 1)
	if err != nil {
		return nil, err
	}

	if req.AvatarUrl != "" {
		var u *url.URL
		u, err = url.Parse(req.AvatarUrl)
		if err != nil {
			return
		}
		u.Host = l.svcCtx.Config.CdnHost
		req.AvatarUrl = u.String()
	}

	_, err = l.svcCtx.UserRPC.UpdateUser(l.ctx, &pb.UpdateUserReq{
		User: &pb.UserDetail{
			Id:        userId,
			AvatarUrl: req.AvatarUrl,
			Nickname:  req.Nickname,
			Motto:     req.Motto,
		},
	})
	if err != nil {
		return nil, err
	}

	return
}
