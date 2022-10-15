package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPrivilegeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPrivilegeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPrivilegeLogic {
	return &UpdateUserPrivilegeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserPrivilegeLogic) UpdateUserPrivilege(req *types.UpdateUserPrivilegeReq) (resp *types.UpdateUserPrivilegeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
