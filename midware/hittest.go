package midware

import (
	"fmt"
	"net/http"
)

func HitTest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsAsset(r.URL.Path) || IsTemplate(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("hit test", r.URL.Path)
			next.ServeHTTP(w, r)
		}
	})
}
