package logic

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCatLogic {
	return &UploadCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  上传或更新猫咪信息
func (l *UploadCatLogic) UploadCat(in *pb.UploadCatReq) (*pb.UploadCatResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UploadCatResp{}, nil
}
