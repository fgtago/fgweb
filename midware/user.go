package midware

import (
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
)

func User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsAsset(r.URL.Path) || IsTemplate(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			ws := appsmodel.GetWebservice()
			ctx := r.Context()
			pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
			pv.IsAuthenticated = ws.Session.GetBool(ctx, string(appsmodel.IsAuthenticatedKeyName))
			if pv.IsAuthenticated {
				pv.UserId = ws.Session.GetString(ctx, string(appsmodel.UserIdKeyName))
				pv.UserNickName = ws.Session.GetString(ctx, string(appsmodel.UserNickNameKeyName))
				pv.UserFullName = ws.Session.GetString(ctx, string(appsmodel.UserFullNameKeyName))
			}
			next.ServeHTTP(w, r)
		}
	})
}
