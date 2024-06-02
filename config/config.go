package config

import (
	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
)

var ws *appsmodel.Webservice
var cfg *appsmodel.Configuration

func New(ws *appsmodel.Webservice) {
	ws = ws
	cfg = &appsmodel.Configuration{}
}

func ReadFromYml(path string) {
	dwlog.Info("read config from %s", path)

}
