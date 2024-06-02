package fgweb

import (
	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/config"
	"github.com/fgtago/fgweb/msg"
)

var ws *appsmodel.Webservice

func New(rootDir string, cfgpath string) (ws *appsmodel.Webservice, err error) {

	// baca configurasi file
	config.New(ws)
	cfg, err := config.ReadFromYml(cfgpath)
	if err != nil {
		dwlog.Error(msg.ErrReadYml, cfgpath)
		return nil, err
	}

	ws = &appsmodel.Webservice{}
	ws.RootDir = rootDir
	ws.Configuration = cfg

	return ws, nil
}

func StartService(port int) (err error) {
	dwlog.Info(msg.InfStartingService, port)

	return nil
}
