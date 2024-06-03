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
}
