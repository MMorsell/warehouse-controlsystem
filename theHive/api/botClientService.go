package botClientService

import (
	"io"
	"log"

	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
)

type Server struct {
}

//Endpoint designated for robot position updated. This information is later relayed to the webclient interface
//TODO: Generate the needed code for the relay of this information to the web interface
func (s *Server) RegisterCurrentPosition(stream botClientService.BotClientService_RegisterCurrentPositionServer) error {
	nrMessages := 0
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Total nr of messages: %d", nrMessages)
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
		log.Printf("Receive position: (%d, %d), message nr %d", point.XPosition, point.YPosition, nrMessages)
	}
}
