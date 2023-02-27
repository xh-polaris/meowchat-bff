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

func toPb(image []types.CreateImageElement) []*pb.CreateImageElement {
	ret := make([]*pb.CreateImageElement, len(image))
	for i := 0; i < len(ret); i++ {
		ret[i] = &pb.CreateImageElement{
			CatId: image[i].CatId,
			Url:   image[i].Url,
		}
	}
	return ret
}

func (l *CreateImageLogic) CreateImage(req *types.CreateImageReq) (resp *types.CreateImageResp, err error) {
	resp = new(types.CreateImageResp)
	data := pb.CreateImageReq{
		Images: toPb(req.Images),
	}
	res, err := l.svcCtx.CollectionRPC.CreateImage(l.ctx, &data)
	if err != nil {
		return
	}
	// 规避错误
	if len(res.ImageIds) > 0 {
		resp.Id = res.ImageIds
	}
	return
}
