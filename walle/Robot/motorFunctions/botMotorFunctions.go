package botMotorFunctions

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"time"

	hiveProto "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protoContract "gits-15.sys.kth.se/Gophers/walle/walle/Robot/proto"
)

/*
Recieve a list of instructions
Move according to these instructions *
before moving, check with hive if it is OK to move
after moving, respond back to hive with the new location
Completing instructions, report back to hive as available
*/

type Robot struct {
	IsAvailable bool
	PositionX   int
	PositionY   int
	Id          string
	PortNumber  int
}

type Coordinate struct {
	xCoordinate int
	yCoordinate int
}

func newCoordinate(CoordinateX int, CoordinateY int) *Coordinate {
	Coordinate := new(Coordinate)
	Coordinate.xCoordinate = CoordinateX
	Coordinate.yCoordinate = CoordinateY
	return Coordinate
}

func (robot *Robot) move(coordinate *Coordinate) {
	robot.PositionX = coordinate.xCoordinate
	robot.PositionY = coordinate.yCoordinate
	// robot.PositionX += coordinate.xCoordinate
	// robot.PositionY += coordinate.yCoordinate
	time.Sleep(time.Second)
}

func (r *Robot) HandleInstructions(instructions *protoContract.Instructions) {
	//Connect to hive
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9738", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect on port 9738: %s", err)
	}
	defer connection.Close()

	//Generate client to hive
	client := hiveProto.NewBotClientServiceClient(connection)

	stream, err := client.RegisterCurrentPosition(context.Background()) //Start position
	if err != nil {
		log.Fatalf("RegisterCurrentPosition failed: %v", err)
	}

	//Update the hive
	for i := range instructions.XMove {
		r.move(&Coordinate{xCoordinate: int(instructions.XMove[i]), yCoordinate: int(instructions.YMove[i])})
		log.Print("Sending update of position to the hive")
		err = stream.Send(&hiveProto.Point{RobotId: r.Id, XPosition: instructions.XMove[i], YPosition: instructions.YMove[i]})

		if err != nil {
			log.Fatalf("Sending update to hive resulted in error: %s", err)
		}
	}

	//Send end call
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}

	if !reply.Recieved {
		log.Fatal("Expected returned bool to be true")
	}

	//Re-register robot as avaliable
	r.reRegisterRobot()

}

func (r *Robot) reRegisterRobot() {
	//Connect to hive
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9738", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect on port 9738: %s", err)
	}
	defer connection.Close()

	//Generate client to hive
	client := hiveProto.NewBotClientServiceClient(connection)

	robotRequestPayload := hiveProto.RegisterRobotPayload{
		XPosition:            int32(r.PositionY),
		YPosition:            int32(r.PositionY),
		RobotEndpointAddress: fmt.Sprint(r.PortNumber)}

	//Register local robot to hive
	response, err := client.RegisterRobot(context.Background(), &robotRequestPayload)
	if err != nil {
		log.Fatalf("Error when calling RegisterRobot: %s", err)
	}

	log.Printf("Success! Setting robot ID to: %s", response.RobotId)

	r.Id = response.RobotId
}

func GetRandomSpawnPoint() (int, int) {
	//Generate local robot
	var RobotSpawns [76]*Coordinate
	for i := 0; i < 20; i++ {
		coordinate := newCoordinate(0, i)
		RobotSpawns[i] = coordinate
	}
	for i := 0; i < 20; i++ {
		coordinate := newCoordinate(19, i)
		RobotSpawns[i+20] = coordinate
	}
	for i := 0; i < 18; i++ {
		coordinate := newCoordinate(i+1, 0)
		RobotSpawns[i+40] = coordinate
	}
	for i := 0; i < 18; i++ {
		coordinate := newCoordinate(i+1, 19)
		RobotSpawns[i+58] = coordinate
	}
	spawnPoint := rand.Intn(len(RobotSpawns))

	return RobotSpawns[spawnPoint].xCoordinate, RobotSpawns[spawnPoint].yCoordinate
}
