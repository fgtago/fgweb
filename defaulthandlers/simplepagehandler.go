package defaulthandlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
)

func SimplePageHandler(pagename string, pv *appsmodel.PageVariable, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ws := appsmodel.GetWebservice()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	// TODO: implmentasikan tpl
	tpl, exists, err := ws.TplMgr.GetPage(pagename, device.Type)
	if err != nil {
		if ws.ShowServerError {
			ErrorPageHandler(500, err.Error(), pv, w, r)
		} else {
			dwlog.Error(err.Error())
			ErrorPageHandler(500, "internal server error", pv, w, r)
		}
		return
	}

	if !exists {
		// error 404
		ErrorPageHandler(404, fmt.Sprintf("page %s not found", pagename), pv, w, r)
		return
	}

	// render page
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, &pv)
	if err != nil {
		if ws.ShowServerError {
			ErrorPageHandler(500, err.Error(), pv, w, r)
		} else {
			dwlog.Error(err.Error())
			ErrorPageHandler(500, "internal server error", pv, w, r)
		}
		return
	}

	// send bufer to browser
	_, err = buff.WriteTo(w)
	if err != nil {
		if ws.ShowServerError {
			ErrorPageHandler(500, err.Error(), pv, w, r)
		} else {
			dwlog.Error(err.Error())
			ErrorPageHandler(500, "internal server error", pv, w, r)
		}
		return
	}
}
