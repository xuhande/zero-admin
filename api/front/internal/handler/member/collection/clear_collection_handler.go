package collection

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/member/collection"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClearCollectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := collection.NewClearCollectionLogic(r.Context(), svcCtx)
		resp, err := l.ClearCollection()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
