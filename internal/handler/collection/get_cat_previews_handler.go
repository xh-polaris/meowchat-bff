package collection

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/collection"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCatPreviewsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := collection.NewGetCatPreviewsLogic(r.Context(), svcCtx)
		resp, err := l.GetCatPreviews()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
