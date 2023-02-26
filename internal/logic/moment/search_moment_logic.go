package moment

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"
	pb3 "github.com/xh-polaris/meowchat-system-rpc/pb"
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
	var data *pb.ListMomentResp
	if req.IsParent {
		data, err = l.svcCtx.MomentRPC.SearchMomentByCommunityId(l.ctx, &pb.SearchMomentByCommunityIdReq{
			CommunityId: req.CommunityId,
			Count:       pageSize,
			Skip:        req.Page * pageSize,
			Keyword:     req.Keyword,
		})
		if err != nil {
			return nil, err
		}
	} else {
		data1, err := l.svcCtx.SystemRPC.ListCommunity(l.ctx, &pb3.ListCommunityReq{
			ParentId: req.CommunityId,
			Size:     -1,
		})
		if err != nil {
			return nil, err
		}

		communityIds := make([]string, 0, len(data1.Communities))
		for i, community := range data1.Communities {
			communityIds[i] = community.Id
		}

		data, err = l.svcCtx.MomentRPC.SearchMomentByMultiCommunityId(l.ctx, &pb.SearchMomentByMultiCommunityIdReq{
			CommunityIds: communityIds,
			Count:        pageSize,
			Skip:         req.Page * pageSize,
			Keyword:      req.Keyword,
		})
		if err != nil {
			return nil, err
		}
	}

	resp.Total = data.Total
	resp.Moments = make([]types.Moment, 0, pageSize)
	err = copier.Copy(&resp.Moments, data.Moments)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data.Moments); i++ {
		user, err := l.svcCtx.UserRPC.GetUser(l.ctx, &pb2.GetUserReq{UserId: data.Moments[i].UserId})
		if err == nil {
			resp.Moments[i].User = types.UserPreview{
				Id:        user.UserId,
				Nickname:  user.Nickname,
				AvatarUrl: user.AvatarUrl,
			}
		}
	}
	return
}
