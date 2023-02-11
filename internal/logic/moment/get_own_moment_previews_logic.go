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

type GetOwnMomentPreviewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOwnMomentPreviewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOwnMomentPreviewsLogic {
	return &GetOwnMomentPreviewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOwnMomentPreviewsLogic) GetOwnMomentPreviews(req *types.GetOwnMomentPreviewsReq) (resp *types.GetOwnMomentPreviewsResp, err error) {
	resp = new(types.GetOwnMomentPreviewsResp)
	userId := l.ctx.Value("userId").(string)
	data, err := l.svcCtx.MomentRPC.ListMomentByUserId(l.ctx, &pb.ListMomentByUserIdReq{
		UserId: userId,
		Count:  pageSize,
		Skip:   req.Page * pageSize,
	})
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
				Id:        user.UserId,
				Nickname:  user.Nickname,
				AvatarUrl: user.AvatarUrl,
			}
		}
	}
	return
}
