package proto

import (
	protoContract "gits-15.sys.kth.se/Gophers/walle/walle/Robot/proto"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) ReceiveTask(ctx context.Context, Instructions *protoContract.Instructions) (*protoContract.HasReceivedTask, error) {
	log.Printf("Recieved Instructions from The Hive: %s", Instructions)
	return &protoContract.HasReceivedTask{Confirmation: "Instructions gathered"}, nil
}
