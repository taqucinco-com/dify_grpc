package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.com/myapp/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedDifyGrpcServiceServer
}

func (s *server) SendMessage(ctx context.Context, in *pb.DifyRequest) (*pb.DifyResponse, error) {
	log.Printf("Received SayHello: %v", in.GetName())
	return &pb.DifyResponse{Answer: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDifyGrpcServiceServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
