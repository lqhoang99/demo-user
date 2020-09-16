package grpcnode

import (
	"demo-user/util"
	"errors"

	"go.mongodb.org/mongo-driver/bson"

	userpb "demo-user/proto/models/user"
	"demo-user/dao"
)

func getUserBriefByID(userIDString string) (*userpb.GetUserBriefByIDResponse, error) {
	var (
		userID,_ = util.HelperParseStringToObjectID(userIDString)
	)

	// Find User
	user, err := dao.UserFindByID(userID)
	if err != nil {
		err = errors.New("Not Found User by ID")
		return nil, err
	}

	// Success
	result := &userpb.GetUserBriefByIDResponse{
		UserBrief: &userpb.UserBrief{
			Id:               userIDString,
			Name:             user.Name,
			TotalTransaction: user.TotalTransaction,
			TotalCommission:  user.TotalCommission,
		},
	}
	return result, nil
}

func updateUserStatsByID(userIDString string, totalTransaction int64, totalCommission float64) (*userpb.UpdateUserStatsByIDResponse, error) {
	var (
		userID,_ = util.HelperParseStringToObjectID(userIDString)
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
		return nil, err
	}

	// Success
	result := &userpb.UpdateUserStatsByIDResponse{}
	return result, nil
}
