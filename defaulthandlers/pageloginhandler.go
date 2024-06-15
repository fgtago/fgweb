package defaulthandlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
)

type LoginData struct {
	LoginError bool
}

// PageLoginHandler handles the HTTP request for the login page.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
// It does not return anything.
func PageLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.Title = fmt.Sprintf("Login - %s", pv.Title)
	pv.Data = &LoginData{}

	serveLoginPage(w, r, pv)
}

// DoLoginHandler handles the HTTP request for the login page.
//
// It takes in an http.ResponseWriter and an http.Request as parameters.
// It does not return anything.
//
// The function retrieves the PageVariable from the context and sets its title to "Login - {original title}".
// It then creates a new LoginData struct and assigns it to the PageVariable's Data field.
// The function retrieves the email and password from the request's form values.
// It logs the email and password using the dwlog.Info function.
// The function sets the LoginError field of the LoginData struct to true.
// The function calls the serveLoginPage function with the provided parameters.
func DoLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.Title = fmt.Sprintf("Login - %s", pv.Title)
	ld := &LoginData{}

	email := r.FormValue("email")
	password := r.FormValue("password")

	dwlog.Info("email: %s, password: %s", email, password)

	// cek login
	ld.LoginError = true

	// assign page data
	pv.Data = ld

	// kalau error, munculin kembali halaman login
	serveLoginPage(w, r, pv)
}

// serveLoginPage renders the login page and sends it to the browser.
//
// Parameters:
// - w: the http.ResponseWriter to write the rendered page to.
// - r: the http.Request that triggered the login page.
// - pv: the *appsmodel.PageVariable containing the data for the login page.
//
// Returns: nothing.
func serveLoginPage(w http.ResponseWriter, r *http.Request, pv *appsmodel.PageVariable) {
	ws := appsmodel.GetWebservice()

	ctx := r.Context()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	tpl, exists, err := ws.TplMgr.GetPage("login", device.Type)
	if err != nil {
		// error 500
		w.WriteHeader(500)
		fmt.Println(err.Error())
	}

	if !exists {
		// error 404
		w.WriteHeader(404)
		fmt.Println("404")
	}

	// render page
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, &pv)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}

	// send bufer to browser
	_, err = buff.WriteTo(w)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}
}
