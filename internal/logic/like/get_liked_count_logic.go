package like

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikedCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLikedCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikedCountLogic {
	return &GetLikedCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLikedCountLogic) GetLikedCount(req *types.GetLikedCountReq) (resp *types.GetLikedCountResp, err error) {
	resp = new(types.GetLikedCountResp)

	likes, err := l.svcCtx.LikeRPC.GetTargetLikes(l.ctx, &pb.GetTargetLikesReq{
		TargetId: req.TargetId,
		Type:     req.TargetType,
	})
	if err != nil {
		return nil, err
	}

	resp.Count = likes.Count

	return
}
