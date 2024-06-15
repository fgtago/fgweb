package main

import (
	"fmt"
	"net/http"

	"github.com/agungdhewe/dwtpl"
	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/go-chi/chi/v5"
)

type PageData struct {
	Nama string
}

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/favicon.ico", defaulthandlers.FaviconHandler)
	fgweb.Get(mux, "/asset/*", defaulthandlers.AssetHandler)
	fgweb.Get(mux, "/template/*", defaulthandlers.TemplateHandler)

	fgweb.Get(mux, "/", Home)
	fgweb.Get(mux, "/about", About)

	return nil

}

func Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "home"
	pv.Data = &PageData{}

	pv.Use(func(pv *appsmodel.PageVariable, layout *dwtpl.Layout) error {
		fmt.Println("on reading page")

		data := pv.Data.(*PageData)
		data.Nama = "jojon"

		return nil
	})

	defaulthandlers.SimplePageHandler(pv, w, r)
}

func About(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "about"
	defaulthandlers.SimplePageHandler(pv, w, r)
}
