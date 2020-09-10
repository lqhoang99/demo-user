package grpcuser

import (
	"log"
	"context"
	"time"

	"demo-user/models"
	transactionpb "demo-transaction/proto/models/transaction"
)
// GetTransactionDetailByUserID ...
func GetTransactionDetailByUserID(userID string) (transactions []models.TransactionDetail, err error) {
	clientConn, client := CreateClient()
	defer clientConn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Call
	result, err := client.GetTransactionDetailByUserID(ctx, &transactionpb.GetTransactionDetailByUserIDRequest{UserID: userID})
	if err != nil {
		log.Printf("Call grpc get transaction by userID error %v\n", err)
		return 
	}
	log.Println("result:",result)

	// Convert to user brief
	transactions = convertToTransactionDetailList(result.TransactionDetail)
	return
}