package defaulthandlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
)

func SimplePageHandler(pv *appsmodel.PageVariable, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ws := appsmodel.GetWebservice()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	// TODO: implmentasikan tpl
	page, err := ws.TplMgr.GetPage(pv.PageName)
	if err != nil {
		if ws.ShowServerError {
			ErrorPageHandler(500, err.Error(), pv, w, r)
		} else {
			dwlog.Error(err.Error())
			ErrorPageHandler(500, "internal server error", pv, w, r)
		}
		return
	}

	// Sesuaikan title
	if page.Config.Title != "" {
		pv.Title = fmt.Sprintf("%s - %s", page.Config.Title, ws.Configuration.Title)
	}

	// jalankan semua middleware
	if pv.MidleWares != nil {
		for _, mw := range *pv.MidleWares {
			err := mw(pv, page.Config)
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
	}

	tpl, inmap := page.Data[device.Type]
	if !inmap {
		// error 404
		ErrorPageHandler(500, fmt.Sprintf("correnponding page template %s not found", pv.PageName), pv, w, r)
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
