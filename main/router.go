package main

import (
	"github.com/fgtago/fgweb"
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
	fgweb.Get(mux, "/login", defaulthandlers.PageLoginHandler)
	fgweb.Post(mux, "/login", defaulthandlers.DoLoginHandler)

	fgweb.Get(mux, "/", defaulthandlers.PageHomeHandler)

	return nil

}
