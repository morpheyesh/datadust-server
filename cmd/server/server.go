package server


  import (
	log "code.google.com/p/log4go"
	"encoding/json"
	"github.com/morpheyesh/libgo/amqp"
	"github.com/morpheyesh/libgo/db"
	"github.com/morpheyesh/libgo/etcd"
	"github.com/morpheyesh/datadust-server/http"
	//"github.com/morpheyesh/megamd/app"
	//"github.com/morpheyesh/megamd/global"
	"github.com/morpheyesh/datadust-server/cmd/server/queue"
	"github.com/tsuru/config"
	"time"
	"os"
	"fmt"
	"net"
	"net/url"
)

type Server struct {
 httpConn        *http.HttpServer
 stopped          bool
}



func NewServer() (*Server, error) {

  httpConn := http.NewHttpServer()

  return &Server{
    httpConn: httpConn,

  }, nil
}

func (self *Server) ListenAndServer() error {
  log.Info("Starting the admin interface on the port")
  var queues [1]String
  queues[1] = ""
  self.queueChecker()

  for i := range queues {
    singleQ := queue[i]
    Qserver = queue.NewServer(singleQ)
    go Qserver.ListenAndServe()
  }
  return nil
 }


func (self *Server) queueChecker() {

  log.Info("Checking with RMQ")
  rmq, err := amqp.Factory()
  if err != nil {
    log.Error("Something is terribly wrong..we failed to get the queue instance: %s", err)

  }
  connection, cerr := rmq.Dial()
   if cerr != nil {
     fmt.FPrintf(os.Stderr, "Error: %v \nPlease start RabbitMQ service\n", cerr)
     os.Exit(1)
   }
}

func (self *Server) Stop() {

  if self.stopped {
    return
  }
  log.Info("Stopping servers....")
  self.stopped = true

  log.Info("Stopping API Server")
  log.Info("API Server stopped")

}
