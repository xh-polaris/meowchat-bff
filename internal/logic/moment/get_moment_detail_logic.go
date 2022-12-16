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

type GetMomentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMomentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMomentDetailLogic {
	return &GetMomentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMomentDetailLogic) GetMomentDetail(req *types.GetMomentDetailReq) (resp *types.GetMomentDetailResp, err error) {
	resp = new(types.GetMomentDetailResp)
	data, err := l.svcCtx.MomentRPC.RetrieveMoment(l.ctx, &pb.RetrieveMomentReq{MomentId: req.MomentId})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp.Moment, data.Moment)
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.UserRPC.GetUser(l.ctx, &pb2.GetUserReq{UserId: data.Moment.UserId})
	if err != nil {
		return nil, err
	}

	resp.Moment.User = types.UserPreview{
		Id:        user.UserId,
		Nickname:  user.Nickname,
		AvatarUrl: user.AvatarUrl,
	}
	return
}
