package grpcnode

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	"demo-user/dao"
	userpb "demo-user/proto/models/user"
	"demo-user/util"
)

func getUserBriefByID(userIDString string) (*userpb.UserBrief, error) {
	var (
		userID = util.HelperParseStringToObjectID(userIDString)
	)

	// Find User
	user, err := dao.UserFindByID(userID)
	if err != nil {
		err = errors.New("Not found user by ID")
		return nil, err
	}

	// Success
	result := &userpb.UserBrief{
		Id:               userIDString,
		Name:             user.Name,
		TotalTransaction: user.TotalTransaction,
		TotalCommission:  user.TotalCommission,
	}
	return result, nil
}

func updateUserStatsByID(userIDString string, totalTransaction int64, totalCommission float64) error {
	var (
		userID = util.HelperParseStringToObjectID(userIDString)
	)

	// Set filter and update
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{
		"totalTransaction": totalTransaction,
		"totalCommission":  totalCommission,
	}}

	// Update User
	err := dao.UserUpdateByID(filter, update)
	if err != nil {
		err = errors.New("Update userStats error")
		return err
	}
	return nil
}
