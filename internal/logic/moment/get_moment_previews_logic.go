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

var pageSize int64 = 10

type GetMomentPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMomentPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMomentPreviewsLogic {
	return &GetMomentPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMomentPreviewsLogic) GetMomentPreviews(req *types.GetMomentPreviewsReq) (resp *types.GetMomentPreviewsResp, err error) {
	resp = new(types.GetMomentPreviewsResp)
	var data *pb.ListMomentResp
	if req.IsParent {
		data, err = l.svcCtx.MomentRPC.ListMomentByCommunityId(l.ctx, &pb.ListMomentByCommunityIdReq{
			CommunityId: req.CommunityId,
			Count:       pageSize,
			Skip:        req.Page * pageSize,
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

		data, err = l.svcCtx.MomentRPC.ListMomentByMultiCommunityId(l.ctx, &pb.ListMomentByMultiCommunityIdReq{
			CommunityIds: communityIds,
			Count:        pageSize,
			Skip:         req.Page * pageSize,
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
