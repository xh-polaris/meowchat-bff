package sts

import (
	"net/http"

	"github.com/xh-polaris/meowchat-bff/internal/logic/sts"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApplySignedUrlAsCommunityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApplySignedUrlAsCommunityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := sts.NewApplySignedUrlAsCommunityLogic(r.Context(), svcCtx)
		resp, err := l.ApplySignedUrlAsCommunity(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
