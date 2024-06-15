package defaulthandlers

import "github.com/fgtago/fgweb/appsmodel"

var currws *appsmodel.Webservice

// New initializes a new instance of the Webservice struct and assigns it to the global variable currws.
//
// Parameters:
// - w: a pointer to a Webservice struct.
func New(w *appsmodel.Webservice) {
	currws = w
}

// GetWebservice returns the current instance of the Webservice struct.
//
// No parameters.
// Returns a pointer to the Webservice struct.
func GetWebservice() *appsmodel.Webservice {
	return currws
}
