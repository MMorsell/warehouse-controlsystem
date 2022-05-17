package main

import (
	"context"
	Proto "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect on port 100: %s", err)
	}
	defer connection.Close()

	client := Proto.NewBotClientServiceClient(connection)

	var RobotSpawns [76]*Coordinate

	for i := 0; i < 20; i++ {
		coordinate := NewCoordinate(0, i)
		RobotSpawns[i] = coordinate
	}
	for i := 0; i < 20; i++ {
		coordinate := NewCoordinate(19, i)
		RobotSpawns[i+20] = coordinate
	}
	for i := 0; i < 18; i++ {
		coordinate := NewCoordinate(i+1, 0)
		RobotSpawns[i+40] = coordinate
	}
	for i := 0; i < 18; i++ {
		coordinate := NewCoordinate(i+1, 19)
		RobotSpawns[i+58] = coordinate
	}

	spawnPoint := rand.Intn(len(RobotSpawns))

	robot := &Robot{true, RobotSpawns[spawnPoint].xCoordinate, RobotSpawns[spawnPoint].yCoordinate, "-1"}

	GridPositions := Proto.GridPositions{RobotId: string(robot.id), XPosition: int32(robot.PositionY), YPosition: int32(robot.PositionY)}

	response, err := client.RegisterRobot(context.Background(), &GridPositions)
	if err != nil {
		log.Fatalf("Error when calling RegisterRobot: %s", err)
	}

	log.Printf("Response from the Server: %s", response.RobotId)

}
