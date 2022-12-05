package post

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/post"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NewPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := post.NewNewPostLogic(r.Context(), svcCtx)
		resp, err := l.NewPost(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
