package defaulthandlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
)

func PageHomeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ws := appsmodel.GetWebservice()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)

	// TODO: implmentasikan tpl
	tpl, exists, err := ws.TplMgr.GetPage("home", device.Type)
	if err != nil {
		// error 500
		fmt.Println(err.Error())
	}

	if !exists {
		// error 404
		fmt.Println("404")
	}

	// render page
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, &pv)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}

	// send bufer to browser
	_, err = buff.WriteTo(w)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		return
	}
}
