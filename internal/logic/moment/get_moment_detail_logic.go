package moment

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"

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
	return
}
