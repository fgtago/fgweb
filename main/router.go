package main

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/main/apps"
	"github.com/go-chi/chi/v5"
)

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/", pagehandlerHome)

	return nil
}

func pagehandlerHome(w http.ResponseWriter, r *http.Request) {
	app := apps.GetApplication()

	app.Webservice.TplMgr.Ready()

	fmt.Fprintln(w, "home page nya ok", app.Webservice.Configuration.Port)
}
