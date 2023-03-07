package user

import (
	"context"
	"github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserLogic) SearchUser(req *types.SearchUserReq) (resp *types.SearchUserResp, err error) {
	resp = new(types.SearchUserResp)
	const pageSize = 10
	data, err := l.svcCtx.UserRPC.SearchUser(l.ctx, &pb.SearchUserReq{
		Nickname: req.Keyword,
		Skip:     req.Page * pageSize,
		Count:    pageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Users = make([]types.UserPreview, 0, len(data.Users))
	for _, user := range data.Users {
		resp.Users = append(resp.Users, types.UserPreview{
			Id:        user.Id,
			Nickname:  user.Nickname,
			AvatarUrl: user.AvatarUrl,
		})
	}
	return
}
