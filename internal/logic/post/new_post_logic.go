package post

import (
	"context"
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/errorx"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewPostLogic {
	return &NewPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewPostLogic) NewPost(req *types.NewPostReq) (resp *types.NewPostResp, err error) {
	resp = new(types.NewPostResp)
	userId := l.ctx.Value("userId").(string)

	if req.Id == "" {
		res, err := l.svcCtx.PostRPC.CreatePost(l.ctx, &pb.CreatePostReq{
			Title:    req.Title,
			Text:     req.Text,
			CoverUrl: req.CoverUrl,
			Tags:     req.Tags,
			UserId:   userId,
		})
		if err != nil {
			return nil, err
		}
		resp.PostId = res.PostId
	} else {
		oldPost, err := l.svcCtx.PostRPC.RetrievePost(l.ctx, &pb.RetrievePostReq{
			PostId: req.Id,
		})
		if err != nil {
			return nil, err
		}
		if oldPost.Post.UserId != userId {
			return nil, errorx.NewCodeError(http.StatusForbidden, "无权修改该帖子")
		}

		_, err = l.svcCtx.PostRPC.UpdatePost(l.ctx, &pb.UpdatePostReq{
			Id:       req.Id,
			Title:    req.Title,
			Text:     req.Text,
			CoverUrl: req.CoverUrl,
			Tags:     req.Tags,
		})
		if err != nil {
			return nil, err
		}
		resp.PostId = req.Id
	}

	return
}
