// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"zero_demo/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/uc/info",
				Handler: UserInfoHandler(serverCtx),
			},
		},
	)
}
