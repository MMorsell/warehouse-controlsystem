package botClientService_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	botClientService "gits-15.sys.kth.se/Gophers/walle/theHive/api"
	serviceContract "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type test struct {
	nrClients                  int
	nrPoints                   int
	randomDelayBetweenMessages bool
}

func setupServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	serviceContract.RegisterBotClientServiceServer(s, &botClientService.Server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func randomPoint(r *rand.Rand) *serviceContract.Point {
	return &serviceContract.Point{XPosition: r.Int31(), YPosition: r.Int31()}
}

func TestRegisterCurrentPositionSingleClient(t *testing.T) {
	tests := []test{
		{nrClients: 1, nrPoints: 10, randomDelayBetweenMessages: false},
		{nrClients: 1, nrPoints: 5, randomDelayBetweenMessages: false},
		{nrClients: 1, nrPoints: 6, randomDelayBetweenMessages: true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d Clients,%d points per client, randomDelayBetweenMessages: %t", tt.nrClients, tt.nrPoints, tt.randomDelayBetweenMessages)
		t.Run(testname, func(t *testing.T) {
			setupServer()
			ctx := context.Background()
			conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				t.Fatalf("Failed to dial bufnet: %v", err)
			}
			defer conn.Close()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))

			client := serviceContract.NewBotClientServiceClient(conn)
			stream, err := client.RegisterCurrentPosition(ctx)
			if err != nil {
				t.Fatalf("RegisterCurrentPosition failed: %v", err)
			}

			for i := 0; i < tt.nrPoints; i++ {
				stream.Send(randomPoint(r))
				if tt.randomDelayBetweenMessages {
					delay := int(r.Int31n(100)) + 100
					time.Sleep(time.Duration(delay) * time.Millisecond)
				}
			}

			reply, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
			}

			if reply.Recieved != true {
				t.Error("Expected returned bool to be true")
			}
			if reply.NumberOfMessages != int32(tt.nrPoints) {
				t.Errorf("Number of points do not match, sent: %d - recieved: %d", tt.nrPoints, reply.NumberOfMessages)
			}
		})
	}
}
