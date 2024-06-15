package config

import (
	"fmt"
	"os"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
	"gopkg.in/yaml.v3"
)

// ReadFromYml reads the YAML configuration file from the specified path.
//
// Parameters:
// - path: the path to the YAML configuration file.
// Returns:
// - *appsmodel.Configuration: a pointer to the configuration struct.
// - error: an error if there was an issue reading or parsing the file.
func ReadFromYml(path string) (cfg *appsmodel.Configuration, err error) {
	// info: membaca file config
	var filedata []byte
	filedata, err = os.ReadFile(path) // baca file dari path
	if err != nil {
		// err: error membaca file
		dwlog.Error(err.Error())
		return nil, fmt.Errorf("cannot read file: %s", path)
	}

	cfg = &appsmodel.Configuration{}
	err = yaml.Unmarshal(filedata, cfg)
	if err != nil {
		// err: error parsing file
		dwlog.Error(err.Error())
		return nil, fmt.Errorf("cannot parse yml file: %s", path)
	}

	return cfg, nil
}
