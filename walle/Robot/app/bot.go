package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	grpcServer "gits-15.sys.kth.se/Gophers/walle/walle/Robot/api"
)

/*
Recieve a list of instructions
Move according to these instructions *
before moving, check with hive if it is OK to move
after moving, respond back to hive with the new location
Completing instructions, report back to hive as available
*/

type Robot struct {
	isAvailable bool
	PositionX   int
	PositionY   int
	id          string
}

type Coordinate struct {
	xCoordinate int
	yCoordinate int
}

//Constructor doesnt seem to work
func NewCoordinate(CoordinateX int, CoordinateY int) *Coordinate {
	Coordinate := new(Coordinate)
	Coordinate.xCoordinate = CoordinateX
	Coordinate.yCoordinate = CoordinateY
	return Coordinate
}

func (robot *Robot) move(coordinate *Coordinate) {
	robot.PositionX += coordinate.xCoordinate
	robot.PositionY += coordinate.yCoordinate
	time.Sleep(time.Second)
}

func (robot *Robot) moveMultiple(listOfCoordinates ...*Coordinate) {
	for i := 0; i < len(listOfCoordinates); i++ {
		robot.move(listOfCoordinates[i])
	}
	robot.isAvailable = true
}

func main() {

	time.Sleep(time.Second * 1)

	//Get random port to use
	portNumber := rand.Intn(1000) + 6000

	//Setup endpoint, but dont listen to it yet
	_, listener, grpcServer := grpcServer.InitServer(portNumber)

	//Init and register the robot as avaliable to the hive
	robot := initRobotAndRegisterAtHive(portNumber)

	fmt.Println(robot)

	//Listen to payloads from the hive
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
