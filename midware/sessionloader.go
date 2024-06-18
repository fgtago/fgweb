package midware

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
)

func SessionLoader(next http.Handler) http.Handler {
	ws := appsmodel.GetWebservice()
	session := ws.Session
	fn := session.LoadAndSave(next)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsAsset(r.URL.Path) || IsTemplate(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			if ws.Configuration.HitTest {
				fmt.Println("handle session", r.URL.Path)
			}
			fn.ServeHTTP(w, r)
		}
	})
}
