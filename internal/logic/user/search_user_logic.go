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
	var pageSize int64 = 10
	if req.Limit != 0 {
		pageSize = req.Limit
	}
	request := &pb.SearchUserReq{
		Nickname: req.Keyword,
		Offset:   new(int64),
		Limit:    new(int64),
	}
	if req.LastToken != "" {
		request.LastToken = &req.LastToken
	}
	*request.Offset = req.Page * pageSize
	*request.Limit = pageSize
	data, err := l.svcCtx.UserRPC.SearchUser(l.ctx, request)
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Token = data.Token
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
