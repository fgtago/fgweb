package defaulthandlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
)

type LoginData struct {
	LoginError bool
}

func PageLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.Title = fmt.Sprintf("Login - %s", pv.Title)
	pv.Data = &LoginData{}

	serveLoginPage(w, r, pv)
}

func DoLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.Title = fmt.Sprintf("Login - %s", pv.Title)
	ld := &LoginData{}

	// cek login
	ld.LoginError = true

	// assign page data
	pv.Data = ld

	// kalau error, munculin kembali halaman login
	serveLoginPage(w, r, pv)
}

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
