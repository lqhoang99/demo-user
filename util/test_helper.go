package util

import (
	"bytes"
	"context"
	"demo-user/config"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"demo-user/models"
	"demo-user/modules/database"
)

var (
	userIDString = "5f24d45125ea51bc57a8285a"
	userID,_       = HelperParseStringToObjectID(userIDString)
	user         = models.UserBSON{
		ID:   userID,
		Name: "Phuc",
	}
)

// HelperToIOReader ...
func HelperToIOReader(i interface{}) io.Reader {
	b, _ := json.Marshal(i)
	return bytes.NewReader(b)
}

// HelperUserCreateFake ..
func HelperUserCreateFake() string {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	//Insert
	_, err := userCol.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}
	return userIDString
}

// HelperConnect ...
func HelperConnect() {
	var (
		envVars = config.GetEnv()
		client  = database.GetClient()
	)

	// Set Database for test ...
	db := client.Database(envVars.Database.TestName)
	fmt.Println("Database Connected to", envVars.Database.TestName)
	database.SetDB(db)
}
