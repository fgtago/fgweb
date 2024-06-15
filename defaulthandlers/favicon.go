package defaulthandlers

import (
	"net/http"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
	"github.com/fgtago/fgweb/appsmodel"
)

// FaviconHandler serves the favicon for the current website.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
// The function retrieves the current website's Webservice, constructs the path to the favicon,
// and serves the favicon file using http.ServeFile.
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	ws := appsmodel.GetWebservice()
	faviconpath := filepath.Join(ws.RootDir, ws.Configuration.Favicon)

	// TODO: cek apakah ada
	exists, _, _ := dwpath.IsFileExists(faviconpath)
	if !exists {
		faviconpath = filepath.Join(ws.CurrentWsDir, "defaulthandlers", "favicon.ico")
	}

	http.ServeFile(w, r, faviconpath)
}
