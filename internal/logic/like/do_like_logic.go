package like

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDoLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoLikeLogic {
	return &DoLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DoLikeLogic) DoLike(req *types.DoLikeReq) (resp *types.DoLikeResp, err error) {
	resp = new(types.DoLikeResp)

	userId := l.ctx.Value("userId").(string)

	_, err = l.svcCtx.LikeRPC.DoLike(l.ctx, &pb.DoLikeReq{
		UserId:       userId,
		TargetId:     req.TargetId,
		Type:         req.TargetType,
		Liked:        req.Liked,
		AssociatedId: "",
	})

	if err != nil {
		return nil, err
	}

	return
}
