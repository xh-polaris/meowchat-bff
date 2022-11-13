package handler

import (
	"net/http"

	"github.com/xh-polaris/cat-community-api/internal/logic"
	"github.com/xh-polaris/cat-community-api/internal/svc"
	"github.com/xh-polaris/cat-community-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCatDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCatDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetCatDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetCatDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
