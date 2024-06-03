package apps

import "github.com/fgtago/fgweb/appsmodel"

type Application struct {
	Webservice *appsmodel.Webservice
}

var app *Application

func New(w *appsmodel.Webservice) {
	app = &Application{}
	app.Webservice = w
}

func GetApplication() *Application {
	return app
}
