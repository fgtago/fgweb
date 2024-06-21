package defaulthandlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
	"github.com/agungdhewe/dwtpl"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/go-chi/chi/v5"
)

// TemplateHandler handles template requests.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
// It retrieves the webservice and sets up allowed assets.
// It retrieves the device from the request context and extracts the filename and extension.
// It checks if the asset can be accessed based on its extension. If it cannot be accessed, it writes an error message to the http.ResponseWriter.
// It checks if there is an asset created for mobile or tablet. If there is, it modifies the filename and checks if the asset exists.
// If the asset exists, it checks if the asset exists. If it does not exist, it writes an error message to the http.ResponseWriter.
// It loads the asset and serves it to the client. If there is an error loading the asset, it writes an error message to the http.ResponseWriter.
func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	ws := GetWebservice()
	setupAllowedAsset(ws)

	ctx := r.Context()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	pathparam := chi.URLParam(r, "*")
	filename := filepath.Base(pathparam)
	extension := filepath.Ext(filename)

	// cek apakah asset boleh diakses
	_, allowed := ws.AllowedAsset[extension]
	if !allowed {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Akses ke asset %s tidak diperbolehkan", filename)
		return
	}

	var exist bool
	var path string

	// apakah ada asset yang dibuat untuk mobile/tablet ?
	if device.Type == dwtpl.DeviceMobile || device.Type == dwtpl.DeviceTablet {
		dir := filepath.Dir(pathparam)
		filename = fmt.Sprintf("%s~%s", filename, device.Type)
		path = filepath.Join(ws.RootDir, ws.Configuration.Template.Dir, dir, filename)

		fmt.Println("template mobile/tablet")
		fmt.Println(path)

		exist, _, _ = dwpath.IsFileExists(path)
		if exist {
			pathparam = filepath.Join(dir, filename)
		}
	}

	// cek apakah asset ada
	path = filepath.Join(ws.RootDir, ws.Configuration.Template.Dir, pathparam)
	exist, _, _ = dwpath.IsFileExists(path)
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
