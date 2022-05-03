package main

import (
	"fmt"
	"log"
	"net"

	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/api"
	serviceContract "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	webServer "gits-15.sys.kth.se/Gophers/walle/theHive/web"
	"google.golang.org/grpc"
)

//contains the channel to web browser subscribers
//messages are routed from grpc -> websocket -> browser
var webSubPool []chan serviceContract.GridPositions

func main() {
	fmt.Println("Hello World!")
	go webServer.SetupWebServer(&webSubPool)
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
	s := botClientService.Server{WebSubPool: &webSubPool}

	grpcServer := grpc.NewServer()
	serviceContract.RegisterBotClientServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
