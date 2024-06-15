package defaulthandlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
)

func ErrorPageHandler(errnum int, errmsg string, pv *appsmodel.PageVariable, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ws := appsmodel.GetWebservice()
	device := ctx.Value(appsmodel.DeviceKeyName).(appsmodel.Device)

	pv.Title = "Error"
	pv.HttpErrorNumber = errnum
	pv.HttpErrorMessage = errmsg

	// TODO: implmentasikan tpl
	tpl, exists, err := ws.TplMgr.GetPage("errorpage", device.Type)
	if err != nil {
		dwlog.Error(err.Error())
		simpleErrorPage(errnum, errmsg, w)
		return
	}

	if !exists {
		simpleErrorPage(errnum, errmsg, w)
		return
	}

	// render page
	buff := new(bytes.Buffer)
	err = tpl.Execute(buff, &pv)
	if err != nil {
		dwlog.Error(err.Error())
		simpleErrorPage(errnum, errmsg, w)
		return
	}

	// send bufer to browser
	w.WriteHeader(errnum)
	_, err = buff.WriteTo(w)
	if err != nil {
		dwlog.Error(err.Error())
		simpleErrorPage(errnum, errmsg, w)
		return
	}

}

func simpleErrorPage(errnum int, errmsg string, w http.ResponseWriter) {
	w.WriteHeader(errnum)

	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Error</title>
		<meta name="viewport"
		      content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">	
		<link rel="icon" type="image/x-icon" href="favicon.ico">
	</head>
	<body>
		<h1>Error %d</h1>
		<p>%s</p>
	</body>
	</html>
	`

	fmt.Fprintf(w, html, errnum, errmsg)
}
