package appsmodel

type Application struct {
	Webservice *Webservice
	RootDir    string
}

var app *Application

func New(w *Webservice) {
	app = &Application{}
	app.Webservice = w
}

func GetApplication() *Application {
	return app
}

func SetRootDir(dir string) {
	app.RootDir = dir
}
