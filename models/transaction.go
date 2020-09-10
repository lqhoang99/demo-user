package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// TransactionDetail ...
	TransactionDetail struct {
		ID                       primitive.ObjectID 
		CompanyID                primitive.ObjectID 
		BranchID                 primitive.ObjectID 
		UserID                   primitive.ObjectID 
		Amount                   float64            
		Commission               float64            
		CompanyCashbackPercent   float64            
		CreatedAt                time.Time          
	}
)
