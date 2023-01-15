package main

import (
	"log"
	"net"

	"github.com/akshitsehgal31/users/transfer"
	"google.golang.org/grpc"
	pb "github.com/akshitsehgal31/users/transfer/User.proto"
)

func main(){
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	s.Serve(lis)
}