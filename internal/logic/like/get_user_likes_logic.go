package like

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikesLogic {
	return &GetUserLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLikesLogic) GetUserLikes(req *types.GetUserLikesReq) (resp *types.GetUserLikesResp, err error) {
	resp = new(types.GetUserLikesResp)
	data, err := l.svcCtx.LikeRPC.GetUserLikes(l.ctx, &pb.GetUserLikesReq{
		UserId:     req.UserId,
		TargetType: req.TargetType,
	})
	if err != nil {
		return nil, err
	}
	resp.Likes = make([]types.Like, 0, len(data.Likes))
	for _, like := range data.Likes {
		resp.Likes = append(resp.Likes, types.Like{
			TargetId:     like.TargetId,
			AssociatedId: like.AssociatedId,
		})
	}
	return
}
