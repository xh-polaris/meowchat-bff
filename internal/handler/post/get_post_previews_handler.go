package post

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/post"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPostPreviewsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostPreviewsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := post.NewGetPostPreviewsLogic(r.Context(), svcCtx)
		resp, err := l.GetPostPreviews(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
