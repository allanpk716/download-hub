// Code generated by goctl. DO NOT EDIT.
package handler

import (
	svc2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc2.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetHandler(serverCtx),
			},
		},
	)
}
