package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/agungdhewe/dwpath"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/go-chi/chi/v5"
)

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	pathparam := chi.URLParam(r, "*")
	serveAsset(pathparam, w, r)
}

func serveAsset(pathparam string, w http.ResponseWriter, r *http.Request) {
	ws := GetWebservice()
	setupAllowedAsset(ws)

	filename := filepath.Base(pathparam)
	extension := filepath.Ext(filename)

	// cek apakah asset boleh diakses
	_, allowed := ws.AllowedAsset[extension]
	if !allowed {
		fmt.Fprintf(w, "Akses ke asset %s tidak diperbolehkan", filename)
		return
	}

	// cek apakah asset ada
	path := filepath.Join(ws.RootDir, "..", pathparam)
	exist, _, _ := dwpath.IsFileExists(path)
	if !exist {
		fmt.Fprintf(w, "Asset %s tidak ditemukan", pathparam)
		return
	}

	// muat asset
	// filesize := fileinfo.Size()
	// contenttype := (*fct)[0]
	filedatasource, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(w, "Error memuat asset %s", pathparam)
		return
	}
	defer filedatasource.Close()

	http.ServeFile(w, r, path)

	// w.Header().Add("Content-Type", contenttype)
	//w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	// w.Header().Add("Content-Length", fmt.Sprintf("%d", filesize))
	// w.WriteHeader(http.StatusOK)
	// io.Copy(w, filedatasource)

}

func setupAllowedAsset(ws *appsmodel.Webservice) {
	if ws.AllowedAsset == nil {
		ws.AllowedAsset = make(map[string]*[]string)

		ws.AllowedAsset[".pdf"] = &[]string{"application/pdf"}

		ws.AllowedAsset[".js"] = &[]string{"text/javascript"}
		ws.AllowedAsset[".mjs"] = &[]string{"application/javascript"}

		ws.AllowedAsset[".css"] = &[]string{"text/css"}

		ws.AllowedAsset[".jpg"] = &[]string{"image/jpeg"}
		ws.AllowedAsset[".jpeg"] = &[]string{"image/jpeg"}
		ws.AllowedAsset[".gif"] = &[]string{"image/gif"}
		ws.AllowedAsset[".png"] = &[]string{"image/png"}
		ws.AllowedAsset[".svg"] = &[]string{"image/svg+xml"}
		ws.AllowedAsset[".png"] = &[]string{"image/png"}

		ws.AllowedAsset[".ico"] = &[]string{"image/vnd.microsoft.icon"}

		ws.AllowedAsset[".woff2"] = &[]string{"font/woff2"}
	}
}
