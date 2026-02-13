package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ninestems/go-grpc-example/internal/client/proto/hello/v1"
)

func main() {
	log.Printf("Connecting to server at localhost:50051")
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Sending request with name: leeroy jenkins")
	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "leeroy jenkins",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	bytes, err := json.Marshal(&resp)
	if err != nil {
		log.Fatalf("Could not marshal: %v", err)
	}

	log.Printf("Response: %s", string(bytes))
}
