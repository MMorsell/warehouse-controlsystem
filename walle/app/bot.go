package main

import (
	"fmt"
)

/*
Recieve a list of instructions
Report to hive as unavaliable
Move according to these instructions *
before moving, check with hive if it is OK to move
after moving, respond back to hive with the new location
Completing instructions, report back to hive as avaliable
*/

type Robot struct {
	isAvailable  bool
	isOkayToMove bool
	PositionX    int
	PositionY    int
}

func (robot *Robot) move(x int, y int) {
	errorMessage := "can not move that far in one method call"
	if robot.isOkayToMove && (!((x < 1 || x > 1) || (y < 1 || y > 1))) {
		robot.PositionX += x
		robot.PositionY += y
	} else {
		fmt.Println(errorMessage)
	}
}

func main() {
	robot := &Robot{true, false, 0, 0}
	fmt.Println(robot)
}
