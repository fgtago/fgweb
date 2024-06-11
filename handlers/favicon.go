package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/fgtago/fgweb/appsmodel"
)

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	ws := appsmodel.GetWebservice()
	faviconpath := filepath.Join(ws.RootDir, ws.Configuration.Favicon)
	http.ServeFile(w, r, faviconpath)
}
