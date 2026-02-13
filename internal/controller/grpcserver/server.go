package grpcserver

import (
	"context"
	"log"
	"time"

	pb "github.com/ninestems/go-grpc-example/internal/controller/proto/hello/v1"
	metav1 "github.com/ninestems/go-grpc-example/internal/controller/proto/meta/v1"
)

// Server реализует HelloService
type Server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello реализует метод SayHello
func (s *Server) SayHello(_ context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request from: %s", req.Name)

	return &pb.HelloResponse{
		MetaResponse: &metav1.MetaResponse{
			Meta: &metav1.Meta{
				RequestId: "my-request-id",
				UserAgent: "some user agent",
				SourceIp:  "try find my ip",
				Timestamp: time.Now().Unix(),
				Headers:   map[string]string{"some map key": "some map value"},
			},
			Message: "Hello from meta" + req.Name,
		},
		Greeting: "Hello " + req.Name,
	}, nil
}
