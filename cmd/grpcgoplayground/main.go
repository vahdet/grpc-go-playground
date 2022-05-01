package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/vahdet/grpc-go-playground/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var port = 50051

type server struct {
	pb.PersonServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetPerson(ctx context.Context, in *pb.GetPersonInput) (*pb.Person, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.Person{
		Id:        in.GetId(),
		Name:      "John Doe",
		BirthDate: timestamppb.New(time.Date(1989, time.Month(2), 15, 10, 30, 0, 0, time.UTC)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
