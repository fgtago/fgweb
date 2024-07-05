package midware

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/justinas/nosurf"
)

func Csrf(next http.Handler) http.Handler {
	ws := appsmodel.GetWebservice()
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   ws.Session.Cookie.Secure,
		SameSite: ws.Session.Cookie.SameSite,
		Path:     ws.Session.Cookie.Path,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsAsset(r.URL.Path) || IsTemplate(r.URL.Path) || IsApi(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			if ws.Configuration.HitTest {
				fmt.Println("csrf", r.URL.Path)
			}
			csrfHandler.ServeHTTP(w, r)
		}
	})
}
