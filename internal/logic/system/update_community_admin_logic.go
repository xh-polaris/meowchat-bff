package system

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/constant"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommunityAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCommunityAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommunityAdminLogic {
	return &UpdateCommunityAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommunityAdminLogic) UpdateCommunityAdmin(req *types.UpdateCommunityAdminReq) (resp *types.UpdateCommunityAdminResp, err error) {
	resp = new(types.UpdateCommunityAdminResp)
	data, err := l.svcCtx.SystemRPC.RetrieveUserRole(l.ctx, &pb.RetrieveUserRoleReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	if !req.IsRemove {
		for _, role := range data.Roles {
			if role.Type == constant.RoleCommunityAdmin && role.CommunityId == req.CommunityId {
				return
				// TODO 应当返回错误
			}
		}
		_, err = l.svcCtx.SystemRPC.UpdateUserRole(l.ctx, &pb.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles: append(data.Roles, &pb.Role{
				Type:        constant.RoleCommunityAdmin,
				CommunityId: req.CommunityId,
			}),
		})
		if err != nil {
			return nil, err
		}
	} else {
		roles := make([]*pb.Role, 0, len(data.Roles))
		for _, role := range data.Roles {
			if role.Type != constant.RoleCommunityAdmin || role.CommunityId != req.CommunityId {
				roles = append(roles, role)
			}
		}
		if len(roles) == len(data.Roles) {
			return
			// TODO 应当返回错误
		}
		_, err = l.svcCtx.SystemRPC.UpdateUserRole(l.ctx, &pb.UpdateUserRoleReq{
			UserId: req.UserId,
			Roles:  roles,
		})
		if err != nil {
			return nil, err
		}
	}
	return
}
