package like

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLikedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikedLogic {
	return &GetUserLikedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLikedLogic) GetUserLiked(req *types.GetUserLikedReq) (resp *types.GetUserLikedResp, err error) {
	resp = new(types.GetUserLikedResp)

	userId := l.ctx.Value("userId").(string)
	like, err := l.svcCtx.LikeRPC.GetUserLike(l.ctx, &pb.GetUserLikedReq{
		UserId:   userId,
		TargetId: req.TargetId,
		Type:     req.TargetType,
	})
	if err != nil {
		return nil, err
	}

	resp.Liked = like.Liked

	return
}
