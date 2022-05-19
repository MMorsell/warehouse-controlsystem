package main

import (
	"context"
	"fmt"
	"log"

	hiveProto "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	motorFunctions "gits-15.sys.kth.se/Gophers/walle/walle/Robot/motorFunctions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Calls the hive and registers this robot as "active",
// this call returns a session-id that will become this robots ID for the lifetime of the order
func initRobotAndRegisterAtHive(thisRobotPortNumber int) motorFunctions.Robot {
	thisRobotsEndpointAddress := fmt.Sprintf(":%d", thisRobotPortNumber)

	//Connect to hive
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9738", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect on port 9738: %s", err)
	}
	defer connection.Close()

	//Generate client to hive
	client := hiveProto.NewBotClientServiceClient(connection)

	xStart, yStart := motorFunctions.GetRandomSpawnPoint()
	robot := &motorFunctions.Robot{
		IsAvailable: true,
		PositionX:   xStart,
		PositionY:   yStart,
		Id:          "-1",
		PortNumber:  thisRobotPortNumber}

	robotRequestPayload := hiveProto.RegisterRobotPayload{
		XPosition:            int32(robot.PositionX),
		YPosition:            int32(robot.PositionY),
		RobotEndpointAddress: thisRobotsEndpointAddress}

	//Register local robot to hive
	response, err := client.RegisterRobot(context.Background(), &robotRequestPayload)
	if err != nil {
		log.Fatalf("Error when calling RegisterRobot: %s", err)
	}

	log.Printf("Success! Setting robot ID to: %s", response.RobotId)

	robot.Id = response.RobotId
	return *robot
}
