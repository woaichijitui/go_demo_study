package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero_mall_study/service/product/api/internal/logic"
	"gozero_mall_study/service/product/api/internal/svc"
	"gozero_mall_study/service/product/api/internal/types"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
