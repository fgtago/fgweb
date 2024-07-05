package fgweb

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"github.com/agungdhewe/dwlog"
	"github.com/agungdhewe/dwtpl"
	"github.com/alexedwards/scs/v2"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/config"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/fgtago/fgweb/midware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RouteHandlerFunc func(mux *chi.Mux) error

var ws *appsmodel.Webservice

// New initializes a new instance of the Webservice struct.
//
// Parameters:
// - rootDir: the root directory of the webservice.
// - cfgpath: the path to the configuration file.
//
// Returns:
// - *appsmodel.Webservice: a pointer to the initialized Webservice struct.
// - error: an error if there was a problem reading the configuration file or preparing the template.
func New(rootDir string, cfgpath string) (*appsmodel.Webservice, error) {
	// siapkan core webservice
	ws = &appsmodel.Webservice{}
	ws.RootDir = rootDir

	// current direktori relative ke fgweb.go
	_, filename, _, _ := runtime.Caller(0)
	ws.CurrentWsDir = filepath.Dir(filename)

	// baca configurasi file
	cfg, err := config.ReadFromYml(cfgpath)
	if err != nil {
		dwlog.Error(err.Error())
		return nil, fmt.Errorf("cannot read config file: %s", cfgpath)
	}
	ws.Configuration = cfg

	// set show server error
	ws.ShowServerError = ws.Configuration.ShowServerError

	// siapkan session manager
	session := scs.New()
	session.Lifetime = time.Duration(ws.Configuration.Cookie.LifeTime) * time.Hour
	session.Cookie.Persist = ws.Configuration.Cookie.Persist
	session.Cookie.Secure = ws.Configuration.Cookie.Secure
	session.Cookie.Path = ws.Configuration.Cookie.Path

	/*
		if ws.Configuration.Cookie.SameSite == "lax" {
			session.Cookie.SameSite = http.SameSiteLaxMode
		} else if ws.Configuration.Cookie.SameSite == "strict" {
			session.Cookie.SameSite = http.SameSiteStrictMode
		} else if ws.Configuration.Cookie.SameSite == "none" {
			session.Cookie.SameSite = http.SameSiteNoneMode
		} else {
			session.Cookie.SameSite = http.SameSiteDefaultMode
		}
	*/

	session.Cookie.SameSite = http.SameSiteLaxMode
	fmt.Println(session.Cookie.SameSite, ws.Configuration.Cookie.SameSite)

	ws.Session = session

	// siapkan keperluan lain
	defaulthandlers.New(ws)

	err = PrepareTemplate(ws)
	if err != nil {
		dwlog.Error(err.Error())
		return nil, fmt.Errorf("error while preparing template")
	}

	appsmodel.NewWebservice(ws)

	return ws, nil
}

func CreateServer(port int, hnd RouteHandlerFunc) *http.Server {
	// buat router
	ws.Mux = httpRequestHandler(hnd)

	// siapkan server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: ws.Mux,
	}

	return srv
}

// StartService starts the web service on the specified port and handles incoming requests using the provided RouteHandlerFunc.
//
// Parameters:
// - port: the port number on which the service should be started.
// - hnd: the RouteHandlerFunc that will handle incoming requests.
//
// Returns:
// - err: an error if there was a problem starting the service.
func StartService(port int, hnd RouteHandlerFunc) (err error) {

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
		dwlog.Error(err.Error())
		return fmt.Errorf("cannot start service on port %d", port)
	}

	return nil
}

func CreateRequestHandler(hnd RouteHandlerFunc) *chi.Mux {
	return httpRequestHandler(hnd)
}

// httpRequestHandler creates a new router and sets up middleware before invoking the provided RouteHandlerFunc.
//
// Parameters:
// - hnd: The RouteHandlerFunc that will handle incoming requests.
//
// Returns:
// - *chi.Mux: The configured router.
func httpRequestHandler(hnd RouteHandlerFunc) *chi.Mux {
	// ws := appsmodel.GetWebservice()

	mux := chi.NewRouter()

	// external middleware
	mux.Use(middleware.Recoverer)

	// testing if page is hit
	if ws.Configuration.HitTest {
		mux.Use(midware.HitTest)
	}

	// internal middleware
	mux.Use(midware.Csrf)
	mux.Use(midware.SessionLoader)
	mux.Use(midware.MobileDetect)
	mux.Use(midware.DefaultPageVariable)
	mux.Use(midware.User)

	// handle dari main program
	hnd(mux)

	// kalau halaman tidak ditemukan
	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
		defaulthandlers.ErrorPageHandler(404, "page not found", pv, w, r)
	})

	return mux
}

// Get registers a route handler function for the given pattern in the chi.Mux router.
//
// Parameters:
//   - mux: The chi.Mux router to register the route handler on.
//   - pattern: The URL pattern for the route.
//   - fn: The route handler function that will handle incoming requests.
//
// Returns: None.
func Get(mux *chi.Mux, pattern string, fn http.HandlerFunc) {
	mux.Get(pattern, func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	})
}

// Post registers a route handler function for the given pattern in the chi.Mux router.
//
// Parameters:
// - mux: The chi.Mux router to register the route handler on.
// - pattern: The URL pattern for the route.
// - fn: The route handler function that will handle incoming requests.
func Post(mux *chi.Mux, pattern string, fn http.HandlerFunc) {
	mux.Post(pattern, func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	})
}

// PrepareTemplate prepares the template for the given webservice.
//
// Parameters:
// - ws: The webservice for which the template needs to be prepared.
//
// Returns:
// - error: An error if the template preparation fails.
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
