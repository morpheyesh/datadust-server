package main

import (

  "github.com/astaxie/beego"
  "github.com/morpheyesh/datadust_api/routers/things"
  "log"
  "strconv"
  "time"
  )

func serverRun() {

  log.Printf("DataDust Server  starting at %s", time.Now())
	beegoHandler()
	log.Println("DataDust Server killed |_|.")


}

func beegoHandler() {

 beego.SessionOn = true
 beego.SetStaticPath("/static_source", "static_source")
 beego.DirectoryIndex = true
 datadust := new(inputData.inputdata)
 beego.Router("/devices/:id", datadust, "get:Get")
 port := "8079"
 http_port, _ := strconv.Atoi(port)
 beego.HttpPort = http_port
 beego.Run()


}
