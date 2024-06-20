package appsmodel

import (
	"net/http"

	"github.com/agungdhewe/dwtpl"
)

type PageVariableMidleware func(pv *PageVariable, pg *dwtpl.PageConfig) error

type PageVariable struct {
	PageName         string
	Title            string
	HttpErrorNumber  int
	HttpErrorMessage string
	Form             *Form
	Data             any
	MidleWares       *[]PageVariableMidleware
	CsrfToken        string
	Request          *http.Request
	Response         http.ResponseWriter
	UserId           string
	UserNickName     string
	UserFullName     string
	IsAuthenticated  bool
}

func (pv *PageVariable) Use(mw PageVariableMidleware) {
	if pv.MidleWares == nil {
		pv.MidleWares = &[]PageVariableMidleware{}
	}
	*pv.MidleWares = append(*pv.MidleWares, mw)
}
