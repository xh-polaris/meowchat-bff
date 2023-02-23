package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles(_ *types.GetUserRolesReq) (resp *types.GetUserRolesResp, err error) {
	resp = new(types.GetUserRolesResp)
	data, err := l.svcCtx.SystemRPC.RetrieveUserRole(l.ctx, &pb.RetrieveUserRoleReq{UserId: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}
	for _, role := range data.Roles {
		resp.Roles = append(resp.Roles, types.Role{
			RoleType:    role.Type,
			CommunityId: role.CommunityId,
		})
	}
	return
}
