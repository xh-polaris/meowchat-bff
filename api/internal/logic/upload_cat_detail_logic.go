package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadCatDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCatDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCatDetailLogic {
	return &UploadCatDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCatDetailLogic) UploadCatDetail(req *types.UploadCatReq) (resp *types.UploadCatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
