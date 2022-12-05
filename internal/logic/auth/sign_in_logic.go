package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xh-polaris/auth-rpc/pb"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	rpcResp, err := l.svcCtx.AuthRPC.SignIn(l.ctx, &pb.SignInReq{
		AuthType: req.AuthType,
		AuthId:   req.AuthId,
		Password: req.Password,
		Params:   req.Params,
	})
	if err != nil {
		return
	}

	auth := l.svcCtx.Config.Auth
	resp = new(types.SignInResp)
	resp.AccessToken, resp.AccessExpire, err = generateJwtToken(rpcResp.UserId, auth.AccessSecret, auth.AccessExpire)
	resp.UserId = rpcResp.UserId
	return
}

func generateJwtToken(userId string, secret string, expire int64) (string, int64, error) {
	iat := time.Now().Unix()
	exp := iat + expire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}
