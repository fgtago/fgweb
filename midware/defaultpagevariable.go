package midware

import (
	"context"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
)

func DefaultPageVariable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pv := &appsmodel.PageVariable{
			Title: "Template",
		}

		ctx := context.WithValue(r.Context(), appsmodel.PageVariableKeyName, pv)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
