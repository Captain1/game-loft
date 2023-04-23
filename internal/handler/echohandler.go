package handler

import (
	"net/http"

	"gameloft/internal/logic"
	"gameloft/internal/svc"
	"gameloft/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EchoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEchoLogic(r.Context(), svcCtx)
		resp, err := l.Echo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
