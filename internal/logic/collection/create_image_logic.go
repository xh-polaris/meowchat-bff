package collection

import (
	"context"
	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateImageLogic {
	return &CreateImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func toPb(image []types.Image) []*pb.Image {
	ret := make([]*pb.Image, len(image))
	for i := 0; i < len(ret); i++ {
		ret[i] = &pb.Image{
			CatId: image[i].CatId,
			Url:   image[i].Url,
		}
	}
	return ret
}

func (l *CreateImageLogic) CreateImage(req *types.CreateImageReq) (resp *types.CreateImageResp, err error) {

	data := pb.CreateImageReq{
		Image: toPb(req.Images),
	}
	res, err := l.svcCtx.CollectionRPC.CreateImage(l.ctx, &data)
	if err != nil {
		return
	}
	resp = &types.CreateImageResp{Id: res.ImageId}
	return
}
