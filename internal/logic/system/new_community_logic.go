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

type NewCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewCommunityLogic {
	return &NewCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewCommunityLogic) NewCommunity(req *types.NewCommunityReq) (resp *types.NewCommunityResp, err error) {
	resp = new(types.NewCommunityResp)

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

	if req.Id == "" {
		data, err := l.svcCtx.SystemRPC.CreateCommunity(l.ctx, &pb.CreateCommunityReq{
			Name:     req.Name,
			ParentId: req.ParentId,
		})
		if err != nil {
			return nil, err
		}
		resp.NewId = data.Id
	} else {
		_, err = l.svcCtx.SystemRPC.UpdateCommunity(l.ctx, &pb.UpdateCommunityReq{
			Id:       req.Id,
			Name:     req.Name,
			ParentId: req.ParentId,
		})
		if err != nil {
			return nil, err
		}
	}
	return
}
