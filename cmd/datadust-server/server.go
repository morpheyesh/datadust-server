package main

import (
	log "code.google.com/p/log4go"
	"github.com/morpheyesh/datadust-server/cmd/datadust-server/server"
	"github.com/tsuru/config"
	"runtime"
	"time"
	//"fmt"
)

func serverRun(dry bool) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	version, _ := config.GetString("version")

	log.Info("Starting dd Server %s...", version)

	server, err := server.NewServer()
	if err != nil {
		// sleep for the log to flush
		time.Sleep(time.Second)
		panic(err)
	}

	//if err := startProfiler(server); err != nil {
		//panic(err)
	//}

	err = server.ListenAndServe()
	if err != nil {
		log.Error("ListenAndServe failed: ", err)
	}
}
