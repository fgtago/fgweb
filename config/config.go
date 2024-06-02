package config

import (
	"github.com/fgtago/fgweb/appsmodel"
)

type Configuration struct {
}

var ws *appsmodel.Webservice

func New(ws *appsmodel.Webservice) {
	ws = ws
}

func ReadFromYml(path string) {

}
