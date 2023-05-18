package moment

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/xh-polaris/meowchat-moment-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"net/url"
)

type NewMomentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewMomentLogic {
	return &NewMomentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewMomentLogic) NewMoment(req *types.NewMomentReq) (resp *types.NewMomentResp, err error) {
	resp = new(types.NewMomentResp)
	m := new(pb.Moment)

	for i := 0; i < len(req.Photos); i++ {
		var u *url.URL
		u, err = url.Parse(req.Photos[i])
		if err != nil {
			return
		}
		u.Host = l.svcCtx.Config.CdnHost
		req.Photos[i] = u.String()
	}

	err = copier.Copy(m, req)
	if err != nil {
		return nil, err
	}

	m.UserId = l.ctx.Value("userId").(string)

	if req.Id == "" {
		var data *pb.CreateMomentResp
		data, err = l.svcCtx.MomentRPC.CreateMoment(l.ctx, &pb.CreateMomentReq{Moment: m})
		resp.MomentId = data.MomentId
	} else {
		_, err = l.svcCtx.MomentRPC.UpdateMoment(l.ctx, &pb.UpdateMomentReq{Moment: m})
		resp.MomentId = req.Id
	}

	if err != nil {
		return nil, err
	}

	return
}
