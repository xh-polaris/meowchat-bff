package user

import (
	"context"
	pb2 "github.com/xh-polaris/meowchat-system-rpc/pb"
	"github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserForAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUserForAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserForAdminLogic {
	return &SearchUserForAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserForAdminLogic) SearchUserForAdmin(req *types.SearchUserForAdminReq) (resp *types.SearchUserForAdminResp, err error) {
	resp = new(types.SearchUserForAdminResp)
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
	resp.Users = make([]types.UserWithRole, 0, len(data.Users))
	for _, user := range data.Users {
		u := types.UserWithRole{
			UserPreview: types.UserPreview{
				Id:        user.Id,
				Nickname:  user.Nickname,
				AvatarUrl: user.AvatarUrl,
			},
		}

		u.Roles = make([]types.Role, 0)
		data, err := l.svcCtx.SystemRPC.RetrieveUserRole(l.ctx, &pb2.RetrieveUserRoleReq{UserId: user.Id})
		if err != nil {
			return nil, err
		}
		for _, role := range data.Roles {
			u.Roles = append(u.Roles, types.Role{
				RoleType:    role.Type,
				CommunityId: role.CommunityId,
			})
		}

		resp.Users = append(resp.Users, u)
	}
	return
}
