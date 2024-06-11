package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/handlers"
	"github.com/go-chi/chi/v5"
)

type PageData struct {
	Nama string
}

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/favicon.ico", handlers.FaviconHandler)
	fgweb.Get(mux, "/asset/*", handlers.AssetHandler)
	fgweb.Get(mux, "/template/*", handlers.TemplateHandler)
	fgweb.Get(mux, "/login", handlers.PageLoginHandler)
	fgweb.Post(mux, "/login", handlers.DoLoginHandler)

	fgweb.Get(mux, "/", pagehandlerHome)

	return nil
}

func pagehandlerHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ws := appsmodel.GetWebservice()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	// TODO: implmentasikan tpl
	tpl, exists, err := ws.TplMgr.GetPage("home", device.Type)
	if err != nil {
		// error 500
		fmt.Println(err.Error())
	}

	if !exists {
		// error 404
		fmt.Println("404")
	}

	pv := &appsmodel.PageVariable{
		Data: PageData{
			Nama: "Agung",
		},
	}

	// render page
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, pv)
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
