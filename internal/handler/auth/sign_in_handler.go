package auth

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/auth"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := auth.NewSignInLogic(r.Context(), svcCtx)
		resp, err := l.SignIn(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
