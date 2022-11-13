package handler

import (
	"net/http"

	"github.com/xh-polaris/cat-community-api/internal/logic"
	"github.com/xh-polaris/cat-community-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAdminsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetAdminsLogic(r.Context(), svcCtx)
		resp, err := l.GetAdmins()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
