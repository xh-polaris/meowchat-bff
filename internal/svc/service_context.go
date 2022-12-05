package svc

import (
	"github.com/xh-polaris/auth-rpc/auth"
	"github.com/xh-polaris/meowchat-bff/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	AuthRPC auth.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthRPC: auth.NewAuth(zrpc.MustNewClient(c.AuthRPC)),
	}
}
