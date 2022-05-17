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
	robot := &Robot{true, RobotSpawns[spawnPoint].xCoordinate, RobotSpawns[spawnPoint].yCoordinate, "A"}
	fmt.Println(robot)
	for i := 0; i < len(RobotSpawns); i++ {
		fmt.Println(RobotSpawns[i])
	}
}
