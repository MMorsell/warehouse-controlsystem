package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, error1 := net.Listen("tcp", "10")
	if error1 != nil {
		log.Fatalf("Failed to listen to port 10: %v", error1)
	}

	grpcServer := new(grpc.Server)

	error2 := grpcServer.Serve(listener)

	if error2 != nil {
		log.Fatalf("Failed to provide grpc server over port 10: %v", error2)
	}

}
