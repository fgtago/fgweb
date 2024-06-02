package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
)

var ws *appsmodel.Webservice

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
	var cfgpath string
	if isDev {
		cfgpath = filepath.Join(rootDir, "config-dev.yml")
	} else {
		cfgpath = filepath.Join(rootDir, "config.yml")
	}

	// start jalankan web
	ws, err = fgweb.New(rootDir, cfgpath)
	if err != nil {
		// ada error saat inisiasi webservice, halt
		panic(err.Error())
	}

	// info: memulai service
	fmt.Println("Starting Service ...")
	port := ws.Configuration.Port
	err = fgweb.StartService(port)
	if err != nil {
		// ada error saat service start, halt
		panic(err.Error())
	}

}
