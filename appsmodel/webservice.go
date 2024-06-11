package appsmodel

import (
	"github.com/agungdhewe/dwtpl"
	"github.com/go-chi/chi/v5"
)

type Webservice struct {
	RootDir       string
	Configuration *Configuration
	Mux           *chi.Mux
	TplMgr        *dwtpl.TemplateManager
	AllowedAsset  map[string]*[]string
}

var ws *Webservice

func NewWebservice(w *Webservice) {
	ws = w
}

func GetWebservice() *Webservice {
	return ws
}

func SetRootDir(dir string) {
	ws.RootDir = dir
}
