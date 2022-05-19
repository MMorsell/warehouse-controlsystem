package robotService

import (
	"context"
	"fmt"
	"log"
	"net"

	motorFunctions "gits-15.sys.kth.se/Gophers/walle/walle/Robot/motorFunctions"
	protoContract "gits-15.sys.kth.se/Gophers/walle/walle/Robot/proto"
	"google.golang.org/grpc"
)

type Server struct {
	address string
	Robot   *motorFunctions.Robot
}

func (s *Server) ReceiveTask(ctx context.Context, Instructions *protoContract.Instructions) (*protoContract.HasReceivedTask, error) {
	log.Printf("Recieved Instructions from The Hive: %s", Instructions)

	go s.Robot.HandleInstructions(Instructions)
	return &protoContract.HasReceivedTask{Confirmation: "Instructions recieved"}, nil
}

//Inits the robot's own grpc server endpoint, but you have to listen to the response from the listener
func InitServer(portNumber int) (*Server, net.Listener, *grpc.Server) {
	listener, error1 := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if error1 != nil {
		log.Fatalf("Failed to listen to port 10: %v", error1)
	}

	s := Server{
		address: fmt.Sprintf(":%d", portNumber),
	}

	grpcServer := grpc.NewServer()
	protoContract.RegisterReceiveTaskServiceServer(grpcServer, &s)

	return &s, listener, grpcServer
}
