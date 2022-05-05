package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Recieve a list of instructions
Move according to these instructions *
before moving, check with hive if it is OK to move
after moving, respond back to hive with the new location
Completing instructions, report back to hive as available
*/

type Robot struct {
	isAvailable  bool
	isOkayToMove bool
	PositionX    int
	PositionY    int
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
	errorMessage := "can not move that far in one method call"
	if robot.isOkayToMove {
		robot.PositionX += coordinate.xCoordinate
		robot.PositionY += coordinate.yCoordinate
		time.Sleep(time.Second)
	} else {
		fmt.Println(errorMessage)
	}
}

func (robot *Robot) moveMultiple(listOfCoordinates ...*Coordinate) {
	for i := 0; i < len(listOfCoordinates); i++ {
		robot.move(listOfCoordinates[i])
	}
	robot.isAvailable = true
}

func main() {
	var RobotSpawns [76]*Coordinate
	for i := 0; i < 20; i++ {
		coordinate := NewCoordinate(1, i+1)
		RobotSpawns[i] = coordinate
	}
	for i := 0; i < 20; i++ {
		coordinate := NewCoordinate(20, i+1)
		RobotSpawns[i+20] = coordinate
	}
	for i := 1; i < 20; i++ {
		coordinate := NewCoordinate(i+1, 1)
		RobotSpawns[i+38] = coordinate
	}
	for i := 1; i < 20; i++ {
		coordinate := NewCoordinate(i+1, 20)
		RobotSpawns[i+56] = coordinate
	}
	spawnPoint := rand.Intn(len(RobotSpawns))
	robot := &Robot{true, true, RobotSpawns[spawnPoint].xCoordinate, RobotSpawns[spawnPoint].yCoordinate}
	fmt.Println(robot)
	robot.move(NewCoordinate(1, 0))
}
