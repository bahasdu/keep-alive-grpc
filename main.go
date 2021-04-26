package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"keep-alive-grpc/pb"
	"log"
	"net"
	// replace this with your own project
)

var counter int = 0

func (s *server) SendRequest(cxt context.Context, r *pb.Request) (*pb.Response, error) {
	counter++
	return &pb.Response{Text: string(counter)}, nil
}

func (s *server) mustEmbedUnimplementedMyServiceServer() {
	panic("implement me")
}

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMyServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
