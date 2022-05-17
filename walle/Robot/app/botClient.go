package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	hiveProto "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initRobotAndRegisterAtHive(thisRobotPortNumber int) Robot {
	thisRobotsEndpointAddress := fmt.Sprintf(":%d", thisRobotPortNumber)

	//Connect to hive
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9738", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect on port 8000: %s", err)
	}
	defer connection.Close()

	//Generate client to hive
	client := hiveProto.NewBotClientServiceClient(connection)

	//Generate local robot
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
	robot := &Robot{
		true,
		RobotSpawns[spawnPoint].xCoordinate,
		RobotSpawns[spawnPoint].yCoordinate,
		"-1"}

	robotRequestPayload := hiveProto.RegisterRobotPayload{
		XPosition:            int32(robot.PositionY),
		YPosition:            int32(robot.PositionY),
		RobotEndpointAddress: thisRobotsEndpointAddress}

	//Register local robot to hive
	response, err := client.RegisterRobot(context.Background(), &robotRequestPayload)
	if err != nil {
		log.Fatalf("Error when calling RegisterRobot: %s", err)
	}

	log.Printf("Success! Setting robot ID to: %s", response.RobotId)

	robot.id = response.RobotId
	return *robot
}
