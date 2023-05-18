package system

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminsLogic {
	return &GetAdminsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminsLogic) GetAdmins(req *types.GetAdminsReq) (resp *types.GetAdminsResp, err error) {
	resp = new(types.GetAdminsResp)
	resp.Admins = make([]types.Admin, 0)

	data, err := l.svcCtx.SystemRPC.ListAdmin(l.ctx, &pb.ListAdminReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Admins, &data.Admins)
	if err != nil {
		return nil, err
	}

	return
}
