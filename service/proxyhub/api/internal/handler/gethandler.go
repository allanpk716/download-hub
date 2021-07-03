package handler

import (
	logic2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/logic"
	svc2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/svc"
	types2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetHandler(ctx *svc2.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types2.GetOneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic2.NewGetLogic(r.Context(), ctx)
		resp, err := l.Get(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
