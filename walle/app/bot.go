package main

import (
	"fmt"
	"math/rand"
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

func NewCoordinate(CoordinateX int, CoordinateY int) *Coordinate {
	Coordinate := new(Coordinate)
	Coordinate.xCoordinate = CoordinateX
	Coordinate.yCoordinate = CoordinateY
	return Coordinate
}

func (robot *Robot) move(coordinate Coordinate) {
	errorMessage := "can not move that far in one method call"
	if robot.isOkayToMove && (!((coordinate.xCoordinate < 1 || coordinate.xCoordinate > 1) || (coordinate.yCoordinate < 1 || coordinate.yCoordinate > 1))) {
		robot.PositionX += coordinate.xCoordinate
		robot.PositionY += coordinate.yCoordinate
	} else {
		fmt.Println(errorMessage)
	}
}

func (robot *Robot) moveMultiple(listOfCoordinates ...Coordinate) {
	for i := 0; i < len(listOfCoordinates); i++ {
		robot.move(listOfCoordinates[i])
	}
	robot.isAvailable = true
}

func main() {
	maxX := 100 //To be decided, measurement for width of storage
	maxY := 100 //To be decided, measurement for height of storage
	robot := &Robot{true, false, rand.Intn(maxX), rand.Intn(maxY)}
	fmt.Println(robot)
}
