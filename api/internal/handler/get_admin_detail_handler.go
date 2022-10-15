package handler

import (
	"net/http"

	"github.com/xh-polaris/cat-community-svc/api/internal/logic"
	"github.com/xh-polaris/cat-community-svc/api/internal/svc"
	"github.com/xh-polaris/cat-community-svc/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAdminDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAdminDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetAdminDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetAdminDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
