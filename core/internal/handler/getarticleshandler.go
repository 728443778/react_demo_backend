package handler

import (
	"net/http"

	"ReactDemoBackend/core/internal/logic"
	"ReactDemoBackend/core/internal/svc"
	"ReactDemoBackend/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArticlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetArticlesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetArticlesLogic(r.Context(), svcCtx)
		resp, err := l.GetArticles(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
