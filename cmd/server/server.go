package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ninestems/go-grpc-example/internal/controller/grpcserver"
	pb "github.com/ninestems/go-grpc-example/internal/controller/proto/hello/v1"
)

func main() {
	// Слушаем порт 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Создаем gRPC сервер
	s := grpc.NewServer()

	// Регистрируем наш сервис
	pb.RegisterHelloServiceServer(s, &grpcserver.Server{})

	log.Println("Server starting on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
