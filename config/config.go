package config

import (
	"os"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/msg"
	"gopkg.in/yaml.v3"
)

func New(w *appsmodel.Webservice) {
}

func ReadFromYml(path string) (cfg *appsmodel.Configuration, err error) {
	// info: membaca file config
	dwlog.Info(msg.InfReadConfig, path)

	var filedata []byte
	filedata, err = os.ReadFile(path) // baca file dari path
	if err != nil {
		// err: error membaca file
		dwlog.Error(msg.ErrReadFile, path)
		return nil, err
	}

	cfg = &appsmodel.Configuration{}
	err = yaml.Unmarshal(filedata, cfg)
	if err != nil {
		// err: error parsing file
		dwlog.Error(msg.ErrParsingFile, path)
		return nil, err
	}

	return cfg, nil
}
