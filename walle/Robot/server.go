package main

import (
	RobotService "gits-15.sys.kth.se/Gophers/walle/walle/Robot/api"
	ProtoContract "gits-15.sys.kth.se/Gophers/walle/walle/Robot/proto"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

func main() {
	portNumber := rand.Intn(1000) + 6000
	listener, error1 := net.Listen("tcp", string(portNumber))
	if error1 != nil {
		log.Fatalf("Failed to listen to port 10: %v", error1)
	}
	S := RobotService.Server{}

	grpcServer := grpc.NewServer()

	ProtoContract.RegisterReceiveTaskServiceServer(grpcServer, &S)

	error2 := grpcServer.Serve(listener)

	if error2 != nil {
		log.Fatalf("Failed to provide grpc server over port 10: %v", error2)
	}
}
