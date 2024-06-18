package appsmodel

import (
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
}

func (pv *PageVariable) Use(mw PageVariableMidleware) {
	if pv.MidleWares == nil {
		pv.MidleWares = &[]PageVariableMidleware{}
	}
	*pv.MidleWares = append(*pv.MidleWares, mw)
}
