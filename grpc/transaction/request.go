package grpcuser

import (
	"context"
	"log"
	"time"

	"demo-user/models"
	transactionpb "demo-user/proto/models/transaction"
)

// GetTransactionDetailByUserID ...
func GetTransactionDetailByUserID(userID string) (transactions []models.TransactionDetail, err error) {
	// Setup client
	clientConn, client := CreateClient()
	defer clientConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call GetTransactionDetailByUserID
	result, err := client.GetTransactionDetailByUserID(ctx, &transactionpb.GetTransactionDetailByUserIDRequest{UserID: userID})
	if err != nil {
		log.Printf("call grpc get transaction by userID error %v\n", err)
		return
	}

	// Convert to user brief
	transactions = convertToTransactionDetailList(result.TransactionDetail)
	return
}
