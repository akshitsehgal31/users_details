package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	// "google.golang.org/genproto/googleapis/ads/googleads/v4/enums"

	// "google.golang.org/genproto/googleapis/ads/googleads/v4/enums"
	// "google.golang.org/genproto"
	// "google.golang.org/genproto/googleapis/ads/googleads/v4/enums"

	"google.golang.org/grpc"

	"github.com/akshitsehgal31/users/transfer"
	
)

type server struct {}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.Response, error) {
    // Insert user into DynamoDB here
    return &pb.Response{Message: "User added successfully"}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.Request) (*pb.User, error) {
    // Get user from DynamoDB here
    return &pb.User{Name: "John", Age: 30}, nil
}

