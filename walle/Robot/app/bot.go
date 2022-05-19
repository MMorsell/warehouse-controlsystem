package main

import (
	"log"
	"math/rand"
	"time"

	grpcServer "gits-15.sys.kth.se/Gophers/walle/walle/Robot/api"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Second * 1)

	//Get random port to use
	portNumber := r.Intn(1000) + 6000

	//Setup endpoint, but dont listen to it yet
	s, listener, grpcServer := grpcServer.InitServer(portNumber)

	//Init and register the robot as avaliable to the hive
	robot := initRobotAndRegisterAtHive(portNumber)
	s.Robot = &robot

	log.Printf("Succesful startup of wall-e robot on port %d", portNumber)
	//Listen to payloads from the hive
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
