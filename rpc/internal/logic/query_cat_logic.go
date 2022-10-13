package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCatLogic {
	return &QueryCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  查询猫咪信息
func (l *QueryCatLogic) QueryCat(in *pb.QueryCatReq) (*pb.QueryCatResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryCatResp{}, nil
}
