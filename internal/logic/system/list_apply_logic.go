package system

import (
	"context"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
	pb2 "github.com/xh-polaris/meowchat-user-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListApplyLogic {
	return &ListApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListApplyLogic) ListApply(req *types.ListApplyReq) (resp *types.ListApplyResp, err error) {
	resp = new(types.ListApplyResp)
	apply, err := l.svcCtx.SystemRPC.ListApply(l.ctx, &pb.ListApplyReq{CommunityId: req.CommunityId})
	if err != nil {
		return nil, err
	}
	applyInfo := make([]types.ApplyInfo, 0, len(apply.Apply))
	for _, x := range apply.Apply {
		user, _ := l.svcCtx.UserRPC.GetUser(l.ctx, &pb2.GetUserReq{UserId: x.ApplicantId})
		userPreview := types.UserPreview{
			Id:        user.User.Id,
			Nickname:  user.User.Nickname,
			AvatarUrl: user.User.AvatarUrl,
		}
		applyInfo = append(applyInfo, types.ApplyInfo{
			UserPreview: userPreview,
			ApplyId:     x.ApplyId,
		})
	}
	resp.ApplyInfo = applyInfo
	return
}
