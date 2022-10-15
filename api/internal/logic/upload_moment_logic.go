package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadMomentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadMomentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadMomentLogic {
	return &UploadMomentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadMomentLogic) UploadMoment(req *types.UploadMomentReq) (resp *types.UploadMomentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
