package main

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/main/apps"
	"github.com/go-chi/chi/v5"
)

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/", pagehandlerHome)

	return nil
}

func pagehandlerHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app := apps.GetApplication()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	// TODO: implmentasikan tpl
	tpl, exists, err := app.Webservice.TplMgr.GetPage("homes", device.Type)
	if err != nil {
		// error 500
		fmt.Println(err.Error())
	}

	if !exists {
		// error 404
		fmt.Println("404")
	}

	fmt.Fprintln(w, "home page nya ok", app.Webservice.Configuration.Port)
}
