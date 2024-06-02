package appsmodel

import "github.com/go-chi/chi/v5"

type Webservice struct {
	RootDir       string
	Configuration *Configuration
	Mux           *chi.Mux
}
