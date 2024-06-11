package fgweb

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/agungdhewe/dwlog"
	"github.com/agungdhewe/dwtpl"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/config"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/fgtago/fgweb/midware"
	"github.com/fgtago/fgweb/msg"
	"github.com/go-chi/chi/v5"
)

type RouteHandlerFunc func(mux *chi.Mux) error

var ws *appsmodel.Webservice

func New(rootDir string, cfgpath string) (*appsmodel.Webservice, error) {
	// siapkan core webservice
	ws = &appsmodel.Webservice{}
	ws.RootDir = rootDir

	// baca configurasi file
	config.New(ws)
	cfg, err := config.ReadFromYml(cfgpath)
	if err != nil {
		dwlog.Error(msg.ErrReadYml, cfgpath)
		return nil, err
	}
	ws.Configuration = cfg

	// siapkan keperluan lain
	defaulthandlers.New(ws)

	err = PrepareTemplate(ws)
	if err != nil {
		dwlog.Error(msg.ErrPrepareTemplate)
		return nil, err
	}

	appsmodel.NewWebservice(ws)

	return ws, nil
}

func StartService(port int, hnd RouteHandlerFunc) (err error) {
	dwlog.Info(msg.InfStartingService, port)

	// buat router
	ws.Mux = httpRequestHandler(hnd)

	// siapkan server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: ws.Mux,
	}

	// jalankan server
	err = srv.ListenAndServe()
	if err != nil {
		dwlog.Error(msg.ErrStartService)
		return err
	}

	return nil
}

func httpRequestHandler(hnd RouteHandlerFunc) *chi.Mux {
	mux := chi.NewRouter()

	// middleware
	mux.Use(midware.MobileDetect)

	// handle dari main program
	hnd(mux)

	return mux
}

func Get(mux *chi.Mux, pattern string, fn http.HandlerFunc) {
	mux.Get(pattern, func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	})
}

func Post(mux *chi.Mux, pattern string, fn http.HandlerFunc) {
	mux.Post(pattern, func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	})
}

func PrepareTemplate(ws *appsmodel.Webservice) error {
	// prepare template
	cfgtpl := &dwtpl.Configuration{
		Dir:    filepath.Join(ws.RootDir, ws.Configuration.Template.Dir),
		Cached: ws.Configuration.Template.Cached,
	}

	mgr, err := dwtpl.New(cfgtpl, ws.Configuration.Template.Options...)
	if err != nil {
		return err
	}

	pagedir := filepath.Join(ws.RootDir, ws.Configuration.Application.PageDir)
	err = mgr.CachePages(pagedir)
	if err != nil {
		return err
	}

	ws.TplMgr = mgr

	return nil
}
