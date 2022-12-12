package svc

import (
	"github.com/xh-polaris/auth-rpc/auth"
	"github.com/xh-polaris/meowchat-bff/internal/config"
	"github.com/xh-polaris/meowchat-collection-rpc/collectionrpc"
	"github.com/xh-polaris/meowchat-moment-rpc/momentrpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	AuthRPC       auth.Auth
	CollectionRPC collectionrpc.CollectionRpc
	MomentRPC     momentrpc.MomentRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AuthRPC:       auth.NewAuth(zrpc.MustNewClient(c.AuthRPC)),
		CollectionRPC: collectionrpc.NewCollectionRpc(zrpc.MustNewClient(c.CollectionRPC)),
		MomentRPC:     momentrpc.NewMomentRpc(zrpc.MustNewClient(c.MomentRPC)),
	}
}
