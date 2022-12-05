package comment

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/comment"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NewCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewNewCommentLogic(r.Context(), svcCtx)
		resp, err := l.NewComment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
