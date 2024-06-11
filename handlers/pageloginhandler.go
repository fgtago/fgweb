package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
)

func PageLoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ws := appsmodel.GetWebservice()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	// TODO: implmentasikan tpl
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
	err = tpl.Execute(buff, nil)
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

func DoLoginHandler(w http.ResponseWriter, r *http.Request) {

}
