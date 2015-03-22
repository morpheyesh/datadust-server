package server


  import (
	log "code.google.com/p/log4go"
	//"encoding/json"
	"github.com/morpheyesh/libgo/amqp"
	//"github.com/morpheyesh/libgo/db"
	//"github.com/morpheyesh/datadust-server/libs/etcd"
	"github.com/morpheyesh/datadust-server/http"
	//"github.com/morpheyesh/datadust-server/app"
	//"github.com/morpheyesh/datadust-server/global"
	"github.com/morpheyesh/datadust-server/cmd/server/queue"
	//"github.com/tsuru/config"
	//"time"
	"os"
	"fmt"
	//"net"
	//"net/url"
)

type Server struct {
 httpConn        *http.HttpServer
 QueueServers    []*queue.QueueServer
 stopped          bool
}



func NewServer() (*Server, error) {

  httpConn := http.NewHttpServer()

  return &Server{
    httpConn: httpConn,

  }, nil
}

func (self *Server) ListenAndServe() error {
  log.Info("Starting the admin interface on the port")
  var queues [1]string
  queues[0] = "test"
  self.queueChecker()

  for i := range queues {
    singleQ := queues[i]
    Qserver := queue.NewServer(singleQ)
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
  log.Debug("connection %v", connection)

   if cerr != nil {
     fmt.Fprintf(os.Stderr, "Error: %v \nPlease start RabbitMQ service\n", cerr)
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
