package midware

import (
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/justinas/nosurf"
)

func Csrf(next http.Handler) http.Handler {
	ws := appsmodel.GetWebservice()
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   ws.Configuration.Cookie.Secure,
		SameSite: http.SameSiteStrictMode,
	})

	return csrfHandler
}
