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

	if req.Limit == nil {
		req.Limit = &pageSize
	}
	request := &pb.ListMomentReq{
		FilterOptions: &pb.FilterOptions{
			OnlyUserId:      req.OnlyUserId,
			OnlyCommunityId: req.CommunityId,
		},
		PaginationOptions: &pb.PaginationOptions{
			Offset:    new(int64),
			Limit:     req.Limit,
			Backward:  req.Backward,
			LastToken: req.LastToken,
		},
	}
	*request.PaginationOptions.Offset = *req.Limit * req.Page
	data, err = l.svcCtx.MomentRPC.ListMoment(l.ctx, request)
	if err != nil {
		return nil, err
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
				Id:        user.User.Id,
				Nickname:  user.User.Nickname,
				AvatarUrl: user.User.AvatarUrl,
			}
		}
	}
	return
}
