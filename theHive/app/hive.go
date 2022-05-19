package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	"gits-15.sys.kth.se/Gophers/walle/theHive/pathfinding"

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
	go orderService(&s)
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

func orderService(s *botClientService.Server) {

	batches := 10
	ch := make(chan []Order, batches)

	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			nrOrders := rand.Intn(4-1) + 1
			interval := rand.Intn(21 - 10)
			ch <- generateOrders(nrOrders)
			time.Sleep(time.Second * time.Duration(interval))
		}
	}()

	var lol []pathfinding.Position
	grid := pathfinding.NewGrid(20, 20, lol)

	t := time.Now().Unix()

	var orders []Order

	for {
		if (*s).AvaliableRobots == nil || len(*s.AvaliableRobots) == 0 {
			continue
		}
		orders := append(orders, <-ch...)
		newTime := time.Now().Unix() - t
		for i := len(*s.AvaliableRobots) - 1; i >= 0; i-- {

			if len(orders) == 0 {
				break //Do nothing more if there is not enough orders to be fulfilled
			}

			targetRobot := (*s.AvaliableRobots)[i]

			//Remove robot from AvaliableRobots
			*s.AvaliableRobots = (*s.AvaliableRobots)[:i]

			// Pop order
			order := orders[len(orders)-1]
			orders = orders[:len(orders)-1]

			start := pathfinding.Position{X: int32(targetRobot.XPosition), Y: int32(targetRobot.YPosition), T: int32(newTime)}
			goal := pathfinding.Position{X: int32(order.pos.x), Y: int32(order.pos.y), T: 0}

			//Path to pick up item
			path := pathfinding.FindPath(grid, start, goal)
			pathfinding.ReservePath(grid, path)

			//Path to drop off item
			newStart := path[len(path)-1]

			dropOffPath := pathfinding.FindPath(grid, newStart, pathfinding.Position{X: int32(order.dropOffPos.x), Y: int32(order.dropOffPos.y), T: 0})
			pathfinding.ReservePath(grid, dropOffPath)

			totalPath := append(path, dropOffPath...)

			s.SendInstructionsToTheRobot(targetRobot, totalPath,
				botClientService.Waypoint{
					Pos:        botClientService.XY{X: order.pos.x, Y: order.pos.y},
					DropOffPos: botClientService.XY{X: order.dropOffPos.x, Y: order.dropOffPos.y},
				})
		}
	}
}

// Generates the provided number of orders (XY coordinates).
func generateOrders(nrOrders int) []Order {
	var items []Order
	for len(items) < nrOrders {
		pos := generateRandomPosition(20, 0)
		if len(items) != 0 {
			it := items[len(items)-1]
			// Literally sampling until new point.
			// Technically O(∞) worst case.
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

// Highly sus code that genereates random wall point
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
