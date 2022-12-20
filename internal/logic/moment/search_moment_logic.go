package moment

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"
	pb2 "github.com/xh-polaris/meowchat-user-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMomentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMomentLogic {
	return &SearchMomentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchMomentLogic) SearchMoment(req *types.SearchMomentReq) (resp *types.SearchMomentResp, err error) {
	resp = new(types.SearchMomentResp)
	data, err := l.svcCtx.MomentRPC.SearchMoment(l.ctx, &pb.SearchMomentReq{
		CommunityId: req.CommunityId,
		Keyword:     req.Keyword,
		Count:       pageSize,
		Skip:        req.Page * pageSize,
	})
	if err != nil {
		return nil, err
	}

	resp.Moments = make([]types.Moment, 0, pageSize)
	err = copier.Copy(&resp.Moments, data.Moments)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data.Moments); i++ {
		user, err := l.svcCtx.UserRPC.GetUser(l.ctx, &pb2.GetUserReq{UserId: data.Moments[i].UserId})
		if err != nil {
			return nil, err
		}

		resp.Moments[i].User = types.UserPreview{
			Id:        user.UserId,
			Nickname:  user.Nickname,
			AvatarUrl: user.AvatarUrl,
		}
	}
	return
}
