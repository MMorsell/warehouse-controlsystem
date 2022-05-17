package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/api"
	serviceContract "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	webServer "gits-15.sys.kth.se/Gophers/walle/theHive/web"
	"google.golang.org/grpc"
)

//contains the subs to web browser subscribers
//messages are routed from grpc -> websocket -> browser
var webSubPool []botClientService.WebSub

func main() {
	go webServer.SetupWebServer(&webSubPool)
	go orderService()
	setupGRPCService()

}

func setupGRPCService() {
	log.Printf("Starting GRPC service!")
	portNumber := ":9738"
	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("Failed to listen to %s to listen: %v", portNumber, err)
	}

	//Start custom service contract
	s := botClientService.Server{WebSubPool: &webSubPool}

	grpcServer := grpc.NewServer()
	serviceContract.RegisterBotClientServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

type XY struct {
	x, y int
}

type Order struct {
	pos        XY
	dropOffPos XY
}

func orderService() {

	batches := 10
	ch := make(chan []Order, batches)

	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			nrOrders := rand.Intn(4-1) + 1
			interval := rand.Intn(21-10) + 10
			ch <- generateOrders(nrOrders)
			time.Sleep(time.Second * time.Duration(interval))
		}
	}()

	for {
		orders := <-ch
		log.Println(orders)
	}
}

// Generates the provided number of orders (XY coordinates).
func generateOrders(nrOrders int) []Order {
	var items []Order
	for len(items) < nrOrders {
		pos := generateRandomPosition(20, 0)
		if len(items) != 0 {
			it := items[len(items)-1]
			if it.pos.x == pos.x && it.pos.y == pos.y {
				continue
			}
		}
		dropOffPos := generateRandomWallPosition()
		items = append(items, Order{pos: pos, dropOffPos: dropOffPos})
	}
	return items
}

// Generates a random XY position pair within [min, max).
func generateRandomPosition(max int, min int) XY {
	x := rand.Intn(max-min) + min
	y := rand.Intn(max-min) + min
	return XY{x: x, y: y}
}

func generateRandomWallPosition() XY {
	branch := rand.Intn(4)
	var x int
	var y int
	// Top wall
	if branch == 0 {
		x = rand.Intn(20)
		y = 0

		// Bottom wall
	} else if branch == 1 {
		x = rand.Intn(20)
		y = 19

		// Left wall
	} else if branch == 2 {
		x = 0
		y = rand.Intn(20)

		// Right wall
	} else {
		x = 19
		y = rand.Intn(20)
	}

	return XY{x: x, y: y}
}
