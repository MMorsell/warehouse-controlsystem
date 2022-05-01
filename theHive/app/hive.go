package main

import (
	"fmt"
	"log"
	"net"

	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/api"
	serviceContract "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World!")
	setupGRPCService()
}

func setupGRPCService() {
	log.Printf("Starting GRPC service!")
	portNumber := ":9738"
	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("Failed to listen to %s to listen: %v", portNumber, err)
	}

	//Start custom service contract
	s := botClientService.Server{}
	grpcServer := grpc.NewServer()
	serviceContract.RegisterBotClientServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
