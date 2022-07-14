package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "gorpc.com/student"
)

var (
	port = flag.Int("port", 8000, "server port")
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) SayHello(ctx context.Context, request *pb.StudentRequest) (*pb.MessageResponse, error) {
	fmt.Printf("Received name: %v", request.GetName())
	fmt.Printf("Received nim: %v", request.GetNim())

	return &pb.MessageResponse{Data: request, CreatedAt: time.Now().String()}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Errorf("failed to listen: %v", err)
	}

	serv := grpc.NewServer()
	pb.RegisterStudentServiceServer(serv, &server{})
	fmt.Printf("server listening on: %v", listener.Addr())

	if err := serv.Serve(listener); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
