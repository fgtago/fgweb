package midware

import (
	"context"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
)

// DefaultPageVariable creates a middleware that sets up the default page variable and passes it to the next handler.
//
// It takes in an http.Handler as a parameter and returns an http.Handler.
func DefaultPageVariable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ws := appsmodel.GetWebservice()

		pv := &appsmodel.PageVariable{
			Title: ws.Configuration.Title,
		}

		ctx := context.WithValue(r.Context(), appsmodel.PageVariableKeyName, pv)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
