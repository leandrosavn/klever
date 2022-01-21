package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	"github.com/lfigueiredo82/klever/pkg/services"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedVoteServiceServer
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterVoteServiceServer(s, &services.VoteService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
