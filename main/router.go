package main

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb"
	"github.com/go-chi/chi/v5"
)

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/", pagehandlerHome)

	return nil
}

func pagehandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home page nya")
}
