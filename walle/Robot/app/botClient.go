package main

import (
	Proto "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect on port 100: %s", err)
	}
	defer connection.Close()

	client := Proto.NewBotClientServiceClient(connection)

}
