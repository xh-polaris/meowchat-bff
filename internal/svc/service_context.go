package svc

import (
	"github.com/xh-polaris/auth-rpc/auth"
	"github.com/xh-polaris/meowchat-bff/internal/config"
	"github.com/xh-polaris/meowchat-collection-rpc/collectionrpc"
	"github.com/xh-polaris/meowchat-comment-rpc/commentrpc"
	"github.com/xh-polaris/meowchat-like-rpc/likerpc"
	"github.com/xh-polaris/meowchat-moment-rpc/momentrpc"
	"github.com/xh-polaris/meowchat-post-rpc/postrpc"
	"github.com/xh-polaris/meowchat-system-rpc/systemrpc"
	"github.com/xh-polaris/meowchat-user-rpc/userrpc"
	"github.com/xh-polaris/sts-rpc/stsrpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	AuthRPC       auth.Auth
	CollectionRPC collectionrpc.CollectionRpc
	MomentRPC     momentrpc.MomentRpc
	SystemRPC     systemrpc.SystemRpc
	LikeRPC       likerpc.Likerpc
	UserRPC       userrpc.UserRpc
	StsRPC        stsrpc.StsRpc
	CommentRPC    commentrpc.CommentRpc
	PostRPC       postrpc.PostRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AuthRPC:       auth.NewAuth(zrpc.MustNewClient(c.AuthRPC)),
		CollectionRPC: collectionrpc.NewCollectionRpc(zrpc.MustNewClient(c.CollectionRPC)),
		MomentRPC:     momentrpc.NewMomentRpc(zrpc.MustNewClient(c.MomentRPC)),
		SystemRPC:     systemrpc.NewSystemRpc(zrpc.MustNewClient(c.SystemRPC)),
		LikeRPC:       likerpc.NewLikerpc(zrpc.MustNewClient(c.LikeRPC)),
		UserRPC:       userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRPC)),
		StsRPC:        stsrpc.NewStsRpc(zrpc.MustNewClient(c.StsRPC)),
		CommentRPC:    commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRPC)),
		PostRPC:       postrpc.NewPostRpc(zrpc.MustNewClient(c.PostRPC)),
	}
}
