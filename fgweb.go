package fgweb

import (
	"fmt"
	"net/http"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/config"
	"github.com/fgtago/fgweb/msg"
	"github.com/go-chi/chi/v5"
)

var ws *appsmodel.Webservice

func New(rootDir string, cfgpath string) (*appsmodel.Webservice, error) {

	// baca configurasi file
	config.New(ws)
	cfg, err := config.ReadFromYml(cfgpath)
	if err != nil {
		dwlog.Error(msg.ErrReadYml, cfgpath)
		return nil, err
	}

	ws = &appsmodel.Webservice{}
	ws.RootDir = rootDir
	ws.Configuration = cfg

	return ws, nil
}

func StartService(port int) (err error) {
	dwlog.Info(msg.InfStartingService, port)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: httpHandler(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		dwlog.Error(msg.ErrStartService)
		return err
	}

	return nil
}

func httpHandler() (mux *chi.Mux) {
	return mux
}
