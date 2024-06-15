package appsmodel

import "net/http"

type PageVariableMidleware func(w http.ResponseWriter, r *http.Request) error

type PageVariable struct {
	PageName         string
	Title            string
	HttpErrorNumber  int
	HttpErrorMessage string
	Data             any
	MidleWares       *[]PageVariableMidleware
}

func (pv *PageVariable) Use(mw PageVariableMidleware) {
	if pv.MidleWares == nil {
		pv.MidleWares = &[]PageVariableMidleware{}
	}
	*pv.MidleWares = append(*pv.MidleWares, mw)
}
