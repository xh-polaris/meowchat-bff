package like

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/pb"
	pb2 "github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikedUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLikedUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikedUsersLogic {
	return &GetLikedUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLikedUsersLogic) GetLikedUsers(req *types.GetLikedUsersReq) (resp *types.GetLikedUsersResp, err error) {
	resp = new(types.GetLikedUsersResp)
	data, err := l.svcCtx.LikeRPC.GetLikedUsers(l.ctx, &pb.GetLikedUsersReq{
		TargetId:   req.TargetId,
		TargetType: req.TargetType,
	})
	if err != nil {
		return nil, err
	}
	resp.Users = make([]types.UserPreview, 0, len(data.UserIds))
	for _, userId := range data.UserIds {
		res, err := l.svcCtx.UserRPC.GetUser(l.ctx, &pb2.GetUserReq{UserId: userId})
		if err != nil {
			logx.Error(err)
		}
		resp.Users = append(resp.Users, types.UserPreview{
			Id:        res.User.Id,
			Nickname:  res.User.Nickname,
			AvatarUrl: res.User.AvatarUrl,
		})
	}
	return
}
