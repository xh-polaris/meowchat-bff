package handler

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetCatsLogic(r.Context(), svcCtx)
		resp, err := l.GetCats()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
