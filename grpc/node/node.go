package grpcnode

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"demo-user/config"
	userpb "demo-user/proto/models/user"
)

// Node ...
type Node struct{}

// GetUserBriefByID ...
func (s *Node) GetUserBriefByID(ctx context.Context, req *userpb.GetUserBriefByIDRequest) (*userpb.GetUserBriefByIDResponse, error) {

	// Get user by id
	data, err := getUserBriefByID(req.GetUserID())

	result := &userpb.GetUserBriefByIDResponse{
		UserBrief: data,
	}
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
	err := updateUserStatsByID(userID, totalTransaction, totalCommission)

	result := &userpb.UpdateUserStatsByIDResponse{}
	return result, err
}

// Start ...
func Start() {
	envVars := config.GetEnv()
	userPort := envVars.GRPCPorts.User

	// Create listen
	lis, err := net.Listen("tcp", userPort)
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	// Create service server
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &Node{})

	// Start server
	log.Println(" gRPC server started on port:" + userPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while %v", err)
	}
}
