package botClientService

import (
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
)

type Server struct {
	//Holds all current web instances subscribed to updates
	WebSubPool *[]chan botClientService.GridPositions
}

//Endpoint designated for robot position updated. This information is later relayed to the webclient interface
func (s *Server) RegisterCurrentPosition(stream botClientService.BotClientService_RegisterCurrentPositionServer) error {
	nrMessages := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	botSessionId := strconv.Itoa(int(r.Int31n(100)) + 2)

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

		if err != nil {
			return err
		}
		nrMessages++
		s.sendUpdateToSubscribers(botClientService.GridPositions{RobotId: botSessionId, XPosition: point.XPosition, YPosition: point.YPosition})
	}
}

func (s *Server) sendUpdateToSubscribers(position botClientService.GridPositions) {
	log.Printf("Sending update to clients")
	for _, sub := range *s.WebSubPool {
		sub <- position
	}
}
