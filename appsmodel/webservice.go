package appsmodel

import (
	"github.com/agungdhewe/dwtpl"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type Webservice struct {
	RootDir         string
	Configuration   *Configuration
	Mux             *chi.Mux
	TplMgr          *dwtpl.TemplateManager
	AllowedAsset    map[string]*[]string
	CurrentWsDir    string
	ShowServerError bool
	Session         *scs.SessionManager
	ExtendedConfig  any
}

var ws *Webservice

// NewWebservice initializes a new instance of the Webservice struct and assigns it to the global variable ws.
//
// Parameters:
// - w: a pointer to a Webservice struct.
func NewWebservice(w *Webservice) {
	ws = w
}

// GetWebservice returns the global instance of the Webservice struct.
//
// No parameters.
// Returns a pointer to a Webservice struct.
func GetWebservice() *Webservice {
	return ws
}

// SetRootDir sets the root directory of the Webservice.
//
// Parameters:
//
//	dir (string): The new root directory.
//
// Returns: None.
func SetRootDir(dir string) {
	ws.RootDir = dir
}
