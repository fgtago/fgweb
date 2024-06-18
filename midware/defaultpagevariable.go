package midware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/justinas/nosurf"
)

// DefaultPageVariable creates a middleware that sets up the default page variable and passes it to the next handler.
//
// It takes in an http.Handler as a parameter and returns an http.Handler.
func DefaultPageVariable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsAsset(r.URL.Path) || IsTemplate(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			ws := appsmodel.GetWebservice()

			pv := &appsmodel.PageVariable{
				Title: ws.Configuration.Title,
			}

			pv.CsrfToken = nosurf.Token(r)

			if ws.Configuration.HitTest {
				fmt.Println("create default variable", r.URL.Path)
			}
			ctx := context.WithValue(r.Context(), appsmodel.PageVariableKeyName, pv)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

	})
}
