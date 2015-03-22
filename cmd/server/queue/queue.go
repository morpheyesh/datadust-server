package queue

import (
	log "code.google.com/p/log4go"
	"github.com/morpheyesh/datadust-server/libs/amqp"
)

type QueueServer struct {
	ListenAddress string
	chann         chan []byte
	shutdown      chan bool
}

//interface arguments
func NewServer(listenAddress string) *QueueServer {
	log.Info("Create New queue server")
	self := &QueueServer{}

	self.ListenAddress = listenAddress
	self.shutdown = make(chan bool, 1)
	log.Info(self)
	return self
}

func (self *QueueServer) ListenAndServe() {
	factor, err := amqp.Factory()
	if err != nil {
		log.Error("Failed to get the queue instance: %s", err)
	}

	pubsub, err := factor.Get(self.ListenAddress)
	if err != nil {
		log.Error("Failed to get the queue instance: %s", err)
	}

	msgChan, _ := pubsub.Sub()
	for msg := range msgChan {
		log.Info("Alrighty! lets priint the data")
   fmt.Println(msg)
	}

	log.Info("Handling message %v", msgChan)
	self.chann = msgChan

	//self.Serve()
}
