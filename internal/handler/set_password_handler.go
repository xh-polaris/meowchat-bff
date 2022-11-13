package handler

import (
	"net/http"

	"github.com/xh-polaris/cat-community-api/internal/logic"
	"github.com/xh-polaris/cat-community-api/internal/svc"
	"github.com/xh-polaris/cat-community-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSetPasswordLogic(r.Context(), svcCtx)
		resp, err := l.SetPassword(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}