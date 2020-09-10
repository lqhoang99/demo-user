package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"demo-user/config"
	"demo-user/models"
	"demo-user/modules/database"
	"demo-user/modules/zookeeper"
)

var (
	// UserString for test
	UserString = "5f24d45125ea51bc57a8285a"

	// UserID for test
	UserID, _ = primitive.ObjectIDFromHex(UserString)
	// User for test
	User = models.UserBSON{
		ID:   UserID,
		Name: "Phuc",
	}
)

// HelperConnect ...
func HelperConnect() {
	zookeeper.Connect()
	envVars := config.GetEnv()

	// Connect
	client, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.URI))
	if err != nil {
		log.Println(err)
		log.Fatal("Cannot connect to database:", envVars.Database.URI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	db := client.Database(envVars.Database.TestName)
	fmt.Println("Database Connected to", envVars.Database.TestName)
	database.SetDB(db)
}

// HelperToIOReader ...
func HelperToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// HelperUserCreateFake ..
func HelperUserCreateFake() {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	//Insert
	_, err := userCol.InsertOne(ctx, User)
	if err != nil {
		log.Println(err)
	}
}
