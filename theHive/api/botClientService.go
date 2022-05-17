package botClientService

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/uuid"

	"gits-15.sys.kth.se/Gophers/walle/theHive/pathfinding"
	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	serviceContract "gits-15.sys.kth.se/Gophers/walle/walle/Robot/proto"
	"google.golang.org/grpc"
)

type Server struct {
	//Holds all current web instances subscribed to updates
	WebSubPool      *[]WebSub
	AvaliableRobots *[]RobotConnection
}

//A web GUI instance
type WebSub struct {
	//Relays updates to the web sub
	Channel *chan botClientService.GridPositions
	//Cleanup connection from server if true
	ClosedConnection *bool
}

type RobotConnection struct {
	robotId      string
	robotAddress string
}

func (s *Server) SendInstructionsToTheRobot(moves []pathfinding.Position) error {
	//Get first robot
	targetRobot := (*s.AvaliableRobots)[len(*s.AvaliableRobots)-1]

	//Remove robot from AvaliableRobots
	*s.AvaliableRobots = (*s.AvaliableRobots)[:len(*s.AvaliableRobots)-1]

	//Create grpc client
	var conn *grpc.ClientConn
	time.Sleep(1 * time.Second) //To allow for hive to fully start
	conn, err := grpc.Dial(fmt.Sprintf(":%s", targetRobot.robotAddress), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	client := serviceContract.NewReceiveTaskServiceClient(conn)

	//convert the information
	xMoves := make([]int32, len(moves))
	yMoves := make([]int32, len(moves))
	for i := 0; i < len(moves); i++ {
		xMoves = append(xMoves, int32(moves[i].X))
		yMoves = append(yMoves, int32(moves[i].Y))
	}

	//Send info
	client.ReceiveTask(context.Background(), &serviceContract.Instructions{XMove: xMoves, YMove: yMoves})
	return nil
}

//Endpoint designated for robot position updated. This information is later relayed to the webclient interface
func (s *Server) RegisterCurrentPosition(stream botClientService.BotClientService_RegisterCurrentPositionServer) error {
	nrMessages := 0
	botSessionId := ""

	for {
		point, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Total nr of messages: %d, robot ID %s", nrMessages, botSessionId)
			log.Printf("Closing connection")
			return stream.SendAndClose(&botClientService.MessageRecieved{
				Recieved:         true,
				NumberOfMessages: int32(nrMessages),
			})
		}

		//Set for first message
		if nrMessages == 1 {
			botSessionId = point.RobotId
		}

		if err != nil {
			return err
		}
		nrMessages++
		s.sendUpdateToSubscribers(botClientService.GridPositions{RobotId: botSessionId, XPosition: point.XPosition, YPosition: point.YPosition})
	}
}

//Endpoint designated for reg. online robots, returns the assigned robot ID
func (s *Server) RegisterRobot(ctx context.Context, point *botClientService.GridPositions) (*botClientService.RobotRegistrationSuccess, error) {

	if s.AvaliableRobots == nil {
		*s.AvaliableRobots = make([]RobotConnection, 10)
	}
	//Get new id for robot
	uuidWithHyphenFunc := uuid.New()
	uuid := uuidWithHyphenFunc.String()
	//Register
	robot := RobotConnection{robotId: uuid, robotAddress: ""}
	*s.AvaliableRobots = append(*s.AvaliableRobots, robot)

	//Return ok with response
	return &botClientService.RobotRegistrationSuccess{RobotId: uuid}, nil
}

func (s *Server) sendUpdateToSubscribers(position botClientService.GridPositions) {
	if s.WebSubPool == nil {
		return
	}

	unsubScribe := []int{}
	for i, sub := range *s.WebSubPool {
		if *sub.ClosedConnection {
			unsubScribe = append(unsubScribe, i)
			continue
		}

		*sub.Channel <- position
	}

	s.removeSubscribersFromPool(unsubScribe)
}

func (s *Server) removeSubscribersFromPool(indexOfPeersToRemove []int) {
	if len(indexOfPeersToRemove) == 0 {
		return
	}

	if len(indexOfPeersToRemove) == len(*s.WebSubPool) {
		//Remove all connections
		*s.WebSubPool = []WebSub{}
		return
	}

	for i := 0; i < len(indexOfPeersToRemove); i++ {
		//Faster to replace index with another remaining peer, than removal
		*s.WebSubPool = replace(*s.WebSubPool, indexOfPeersToRemove[i])
	}
}

func replace(s []WebSub, i int) []WebSub {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
