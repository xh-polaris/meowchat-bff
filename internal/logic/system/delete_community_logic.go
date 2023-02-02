package system

import (
	"context"
	"github.com/xh-polaris/meowchat-bff/internal/errorx"
	"github.com/xh-polaris/meowchat-system-rpc/constant"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommunityLogic {
	return &DeleteCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommunityLogic) DeleteCommunity(req *types.DeleteCommunityReq) (resp *types.DeleteCommunityResp, err error) {
	resp = new(types.DeleteCommunityResp)

	userId := l.ctx.Value("userId").(string)
	result, err := l.svcCtx.SystemRPC.ContainsRole(l.ctx, &pb.ContainsRoleReq{
		UserId: userId,
		Role: &pb.Role{
			Type: constant.RoleSuperAdmin,
		},
	})
	if err != nil || !result.Contains {
		err = errorx.NewForbiddenError("您没有权限进行此操作")
		return
	}

	_, err = l.svcCtx.SystemRPC.DeleteCommunity(l.ctx, &pb.DeleteCommunityReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return
}
