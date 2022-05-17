package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	serviceContract "gits-15.sys.kth.se/Gophers/walle/theHive/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	time.Sleep(1 * time.Second) //To allow for hive to fully start
	conn, err := grpc.Dial(":9738", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	client := serviceContract.NewBotClientServiceClient(conn)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 300 // Traverse at least two points
	var points []*serviceContract.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}

	log.Printf("Traversing %d points.", len(points))
	stream, err := client.RegisterCurrentPosition(context.Background())
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
	}
	for _, point := range points {
		time.Sleep(time.Second)
		log.Printf("One more point")
		if err := stream.Send(point); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
}

func randomPoint(r *rand.Rand) *serviceContract.Point {
	return &serviceContract.Point{XPosition: r.Int31(), YPosition: r.Int31()}
}
