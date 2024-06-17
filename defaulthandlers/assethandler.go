package defaulthandlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/go-chi/chi/v5"
)

// AssetHandler handles asset requests.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
func AssetHandler(w http.ResponseWriter, r *http.Request) {
	pathparam := chi.URLParam(r, "*")
	serveAsset(pathparam, w, r)
}

// serveAsset handles asset requests.
//
// It takes in a string representing the path to asset parameter, an http.ResponseWriter, and an http.Request as parameters.
// It checks if the asset can be accessed based on its extension. If it can be accessed, it checks if the asset exists.
// If the asset exists, it loads the asset and serves it to the client. If there is an error loading the asset, it writes an error message to the http.ResponseWriter.
func serveAsset(pathparam string, w http.ResponseWriter, r *http.Request) {
	ws := GetWebservice()
	setupAllowedAsset(ws)

	filename := filepath.Base(pathparam)
	extension := filepath.Ext(filename)

	// cek apakah asset boleh diakses
	_, allowed := ws.AllowedAsset[extension]
	if !allowed {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Akses ke asset %s tidak diperbolehkan", filename)
		return
	}

	// cek apakah asset ada
	path := filepath.Join(ws.RootDir, pathparam)
	exist, _, _ := dwpath.IsFileExists(path)
	if !exist {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Asset %s tidak ditemukan", pathparam)
		return
	}

	// muat asset
	// filesize := fileinfo.Size()
	// contenttype := (*fct)[0]
	filedatasource, err := os.Open(path)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error memuat asset %s", pathparam)
		return
	}
	defer filedatasource.Close()

	http.ServeFile(w, r, path)

}

// setupAllowedAsset initializes the AllowedAsset map in the given Webservice struct.
//
// It checks if the AllowedAsset map is nil and if so, creates a new map. Then, it sets default values for various file extensions and their corresponding MIME types.
// The function does not take any parameters and does not return anything.
func setupAllowedAsset(ws *appsmodel.Webservice) {
	if ws.AllowedAsset == nil {
		ws.AllowedAsset = make(map[string]*[]string)

		ws.AllowedAsset[".pdf"] = &[]string{"application/pdf"}
		ws.AllowedAsset[".mjs"] = &[]string{"application/javascript"}

		ws.AllowedAsset[".css"] = &[]string{"text/css"}
		ws.AllowedAsset[".js"] = &[]string{"text/javascript"}

		ws.AllowedAsset[".jpeg"] = &[]string{"image/jpeg"}
		ws.AllowedAsset[".jpg"] = &[]string{"image/jpeg"}
		ws.AllowedAsset[".gif"] = &[]string{"image/gif"}
		ws.AllowedAsset[".png"] = &[]string{"image/png"}
		ws.AllowedAsset[".svg"] = &[]string{"image/svg+xml"}
		ws.AllowedAsset[".png"] = &[]string{"image/png"}
		ws.AllowedAsset[".ico"] = &[]string{"image/vnd.microsoft.icon"}

		ws.AllowedAsset[".woff2"] = &[]string{"font/woff2"}

	}
}
