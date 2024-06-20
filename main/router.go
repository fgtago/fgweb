package main

import (
	"net/http"

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

	fgweb.Get(mux, "/", DefaultHome)
	fgweb.Get(mux, "/about", DefaultAbout)

	return nil

}

func DefaultHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "home"
	pv.Data = &PageData{}

	// pv.Use(func(w http.ResponseWriter, r *http.Request) error {
	// 	fmt.Println("on reading page")

	// 	pv.Title = "Home - " + pv.Title
	// 	data := pv.Data.(*PageData)
	// 	data.Nama = "jojon"

	// 	return nil
	// })

	defaulthandlers.SimplePageHandler(pv, w, r)
}

func DefaultAbout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "about"
	defaulthandlers.SimplePageHandler(pv, w, r)
}
