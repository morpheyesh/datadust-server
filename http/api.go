package http

import (
  //log "code.google.com/p/log4go"
  "github.com/bmizerany/pat"
  "github.com/tsuru/config"
  "net"
  //libhttp "net/http"
  //"strconv"
  //"strings"
  "time"
)

type HttpServer struct {
	conn        net.Listener
	HttpPort    int
//adminAssets Dir string
	shutdown    chan bool
	readTimeout time.Duration
	p           *pat.PatternServeMux
}

func NewHttpServer() *HttpServer {

  apiHttpPortString, _ := config.GetInt("admin:port")
  self := &HttpServer{}
  self.HttpPort = apiHttpPortString
  self.shutdown = make(chan bool, 2)
  self.p = pat.New()
  self.readTimeout = 10 * time.Second
  return self

}
