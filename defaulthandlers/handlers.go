package defaulthandlers

import "github.com/fgtago/fgweb/appsmodel"

var currws *appsmodel.Webservice

func New(w *appsmodel.Webservice) {
	currws = w
}

func GetWebservice() *appsmodel.Webservice {
	return currws
}
