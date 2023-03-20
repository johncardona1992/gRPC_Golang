package main

import (
	"context"
	"log"
	"time"
	pb "zelator/grpc/greet/proto"
)

func doGreetLongGreet(c pb.GreetServiceClient) {
	log.Println("doGreetLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Zelator"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %s\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
