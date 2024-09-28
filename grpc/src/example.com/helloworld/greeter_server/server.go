package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"strings"

	pb "example.com/myapp/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received SayHello: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

// SayHelloAgain implements helloworld.GreeterServer
func (s *server) SayHelloAgain(in *pb.HelloRequest, stream pb.Greeter_SayHelloAgainServer) error {
	log.Printf("Received SayHelloAgain: %v", in.GetName())
	for i := 0; i < 2; i++ {
		res := &pb.HelloResponse{Message: "Hello " + in.GetName()}
		if err := stream.Send(res); err != nil {
			log.Fatalf("failed to send: %v", err)
			return err
		}
		time.Sleep(3 * time.Second)
	}
	log.Printf("Closed SayHelloAgain")
	return nil
}

func (s *server) SayHelloToMany(stream pb.Greeter_SayHelloToManyServer) error {
	log.Printf("Open SayHelloToMany")
	var slice []string
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			joinedString := strings.Join(slice, ", ")
			log.Printf("Closed SayHelloToMany")
			return stream.SendAndClose(&pb.HelloResponse{
				Message: "Hello! " + joinedString,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("Received SayHelloToMany: %v", in.Name)
		slice = append(slice, in.Name)
		time.Sleep(3 * time.Second)
	}
}

func (s *server) SayChat(stream pb.Greeter_SayChatServer) error {
	log.Printf("Open SayChat")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Closed SayChat")
			return nil
		}
		if err != nil {
			return err
		}
		name := in.Name
		log.Printf("Received SayChat: %v", in.Name)
		time.Sleep(3 * time.Second)
		log.Printf("Send SayChat: %v", in.Name)
		res := &pb.HelloResponse{Message: "Hello " + name}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
