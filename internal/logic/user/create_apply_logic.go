package user

import (
	"context"
	"github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApplyLogic {
	return &CreateApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApplyLogic) CreateApply(req *types.CreateApplyReq) (resp *types.CreateApplyResp, err error) {
	resp = new(types.CreateApplyResp)
	ApplicantId := l.ctx.Value("userId").(string)
	_, err = l.svcCtx.SystemRPC.CreateApply(l.ctx, &pb.CreateApplyReq{
		ApplicantId: ApplicantId,
		CommunityId: req.CommunityId,
	})
	if err != nil {
		return nil, err
	}
	return
}
