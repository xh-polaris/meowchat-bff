package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	CdnHost string
	CIHost  string
	CosApi  struct {
		SecretId  string
		SecretKey string
	}
	AuthRPC       zrpc.RpcClientConf
	CollectionRPC zrpc.RpcClientConf
	MomentRPC     zrpc.RpcClientConf
	SystemRPC     zrpc.RpcClientConf
	LikeRPC       zrpc.RpcClientConf
	UserRPC       zrpc.RpcClientConf
	StsRPC        zrpc.RpcClientConf
	CommentRPC    zrpc.RpcClientConf
	PostRPC       zrpc.RpcClientConf
}
