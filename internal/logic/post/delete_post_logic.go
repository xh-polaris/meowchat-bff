package post

import (
	"context"
	"github.com/xh-polaris/meowchat-bff/internal/errorx"
	"github.com/xh-polaris/meowchat-post-rpc/pb"
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (resp *types.DeletePostResp, err error) {
	resp = new(types.DeletePostResp)
	userId := l.ctx.Value("userId").(string)

	post, err := l.svcCtx.PostRPC.RetrievePost(l.ctx, &pb.RetrievePostReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if post.Post.UserId != userId {
		return nil, errorx.NewCodeError(http.StatusForbidden, "无权删除该帖子")
	}

	_, err = l.svcCtx.PostRPC.DeletePost(l.ctx, &pb.DeletePostReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return
}
