package main

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/handlers"
	"github.com/fgtago/fgweb/main/apps"
	"github.com/go-chi/chi/v5"
)

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/favicon.ico", favicon)
	fgweb.Get(mux, "/asset/*", handlers.AssetHandler)
	fgweb.Get(mux, "/template/*", handlers.TemplateHandler)
	fgweb.Get(mux, "/login", handlers.PageLoginHandler)

	fgweb.Get(mux, "/", pagehandlerHome)

	return nil
}

func favicon(w http.ResponseWriter, r *http.Request) {
	app := apps.GetApplication()
	faviconpath := filepath.Join(app.RootDir, app.Webservice.Configuration.Favicon)
	fmt.Println(faviconpath)
	http.ServeFile(w, r, faviconpath)
}

func pagehandlerHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app := apps.GetApplication()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	// TODO: implmentasikan tpl
	tpl, exists, err := app.Webservice.TplMgr.GetPage("home", device.Type)
	if err != nil {
		// error 500
		fmt.Println(err.Error())
	}

	if !exists {
		// error 404
		fmt.Println("404")
	}

	// render page
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, nil)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}

	// send bufer to browser
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}

	//fmt.Fprintln(w, "home page nya ok", app.Webservice.Configuration.Port)
}
