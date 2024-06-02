package fgweb

import (
	"path/filepath"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/config"
)

var ws *appsmodel.Webservice

func New(rootDir string, configFileName string) (err error) {
	ws = &appsmodel.Webservice{}
	ws.RootDir = rootDir

	// baca configurasi file
	cfgpath := filepath.Join(rootDir, configFileName)
	config.New(ws)
	config.ReadFromYml(cfgpath)

	return nil
}

func StartService() (err error) {

	return nil
}
