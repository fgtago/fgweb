package fgweb

import (
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/config"
)

var ws *appsmodel.Webservice

func New(rootDir string, cfgpath string) (err error) {
	ws = &appsmodel.Webservice{}
	ws.RootDir = rootDir

	// baca configurasi file
	config.New(ws)
	config.ReadFromYml(cfgpath)

	return nil
}

func StartService() (err error) {

	return nil
}
