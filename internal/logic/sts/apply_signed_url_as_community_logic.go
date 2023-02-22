package sts

import (
	"context"
	"github.com/google/uuid"
	"github.com/xh-polaris/sts-rpc/pb"
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplySignedUrlAsCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplySignedUrlAsCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplySignedUrlAsCommunityLogic {
	return &ApplySignedUrlAsCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplySignedUrlAsCommunityLogic) ApplySignedUrlAsCommunity(req *types.ApplySignedUrlAsCommunityReq) (resp *types.ApplySignedUrlAsCommunityResp, err error) {
	resp = new(types.ApplySignedUrlAsCommunityResp)
	data, err := l.svcCtx.StsRPC.GenCosSts(l.ctx, &pb.GenCosStsReq{Path: req.CommunityId + "/*"})
	if err != nil {
		return nil, err
	}
	resp.SessionToken = data.SessionToken
	if req.Prefix != "" {
		req.Prefix += "/"
	}
	data2, err := l.svcCtx.StsRPC.GenSignedUrl(l.ctx, &pb.GenSignedUrlReq{
		SecretId:  data.SecretId,
		SecretKey: data.SecretKey,
		Method:    http.MethodPut,
		Path:      "communities/" + req.CommunityId + "/" + req.Prefix + uuid.New().String() + req.Suffix,
	})
	resp.Url = data2.SignedUrl
	return
}
