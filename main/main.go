package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/fgtago/fgweb"
)

func main() {
	var err error

	// baca parameter dari cli
	var isDev bool
	flag.BoolVar(&isDev, "dev", false, "Jika diset --dev akan menggunakan config mode development")
	flag.Parse()

	// set root direktori ke current working direktori
	// rootDir, err := os.Getwd()
	_, filename, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(filename)

	// ambil file konfigurasi
	var cfgFileName string
	if isDev {
		cfgFileName = filepath.Join(rootDir, "config-dev.yml")
	} else {
		cfgFileName = filepath.Join(rootDir, "config.yml")
	}

	// start jalankan web
	err = fgweb.New(rootDir, cfgFileName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Staring Service...")
	err = fgweb.StartService()
	if err != nil {
		panic(err.Error())
	}
}
