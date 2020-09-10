package grpcnode

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	userpb "demo-user/proto/models/user"
	"demo-user/config"
)

// Node ...
type Node struct{}

// GetUserBriefByID ...
func (s *Node) GetUserBriefByID(ctx context.Context, req *userpb.GetUserBriefByIDRequest) (*userpb.GetUserBriefByIDResponse, error) {
	var (
		userID = req.GetUserID()
	)

	// Get user by id
	result, err := getUserBriefByID(userID)

	return result, err
}

// UpdateUserStatsByID ...
func (s *Node) UpdateUserStatsByID(ctx context.Context, req *userpb.UpdateUserStatsByIDRequest) (*userpb.UpdateUserStatsByIDResponse, error) {
	var (
		userID           = req.GetId()
		totalTransaction = req.GetTotalTransaction()
		totalCommission  = req.GetTotalCommission()
	)

	// Update userStats
	result, err := updateUserStatsByID(userID, totalTransaction, totalCommission)

	return result, err
}

// Start ...
func Start() {
	envVars := config.GetEnv()
	userPort := envVars.GRPCPorts.User

	// Create Listen
	lis, err := net.Listen("tcp", userPort)
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	// Create Service Server
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &Node{})

	log.Println(" gRPC server started on port:" + userPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while %v", err)
	}
}
