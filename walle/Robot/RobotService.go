package Robot

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) askForTask(ctx context.Context, path *Instructions) (*HasReceivedTask, error) {
	log.Printf("Recieved Instructions from The Hive: %s", path.Instruction)
	return &HasReceivedTask{Confirmation: "Instructions gathered"}, nil
}
