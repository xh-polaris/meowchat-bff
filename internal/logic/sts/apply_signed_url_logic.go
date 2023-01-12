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

type ApplySignedUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplySignedUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplySignedUrlLogic {
	return &ApplySignedUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplySignedUrlLogic) ApplySignedUrl(req *types.ApplySignedUrlReq) (resp *types.ApplySignedUrlResp, err error) {
	resp = new(types.ApplySignedUrlResp)
	userId := l.ctx.Value("userId").(string)
	data, err := l.svcCtx.StsRPC.GenCosSts(l.ctx, &pb.GenCosStsReq{UserId: userId, Path: "*"})
	resp.SessionToken = data.SessionToken
	if req.Prefix != "" {
		req.Prefix += "/"
	}
	data2, err := l.svcCtx.StsRPC.GenSignedUrl(l.ctx, &pb.GenSignedUrlReq{
		SecretId:  data.SecretId,
		SecretKey: data.SecretKey,
		Method:    http.MethodPut,
		Path:      "users/" + userId + "/" + req.Prefix + uuid.New().String() + req.Suffix,
	})
	resp.Url = data2.SignedUrl
	return
}
