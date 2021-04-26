package main

import (
	"fmt"
	"github.com/bahasdu/keep-alive-grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync/atomic"
)

var counter uint64

func (s *Server) SendRequest(context.Context, *pb.Request) (*pb.Response, error) {
	atomic.AddUint64(&counter, 1)
	fmt.Println(counter)
	return &pb.Response{Text: fmt.Sprintf("%d", counter)}, nil
}

type Server struct {
	// Embed the unimplemented server
	pb.UnimplementedMyServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	s := Server{}
	pb.RegisterMyServiceServer(grpcServer, &s)

	grpcServer.Serve(lis)
}
