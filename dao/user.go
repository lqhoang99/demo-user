package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-user/models"
	"demo-user/modules/database"
)

// UserList ...
func UserList() ([]models.UserBSON, error) {
	var (
		companyCol = database.UserCol()
		ctx        = context.Background()
		doc        = make([]models.UserBSON, 0)
	)

	// Find
	cursor, err := companyCol.Find(ctx, bson.M{})

	// Close cursor
	defer cursor.Close(ctx)

	// Set result
	cursor.All(ctx, &doc)
	return doc, err
}

// UserCreate ...
func UserCreate(doc models.UserBSON) (models.UserBSON, error) {
	var (
		collection = database.UserCol()
		ctx        = context.Background()
	)

	// Insert one
	_, err := collection.InsertOne(ctx, doc)
	return doc, err
}

// UserFindByID ...
func UserFindByID(id primitive.ObjectID) (models.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		result  models.UserBSON
		filter  = bson.M{"_id": id}
	)

	// Find
	err := userCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// UserUpdateByID ...
func UserUpdateByID(filter bson.M, updateData bson.M) (err error) {
	var (
		userCol = database.UserCol()
		ctx       = context.Background()
	)

	_, err = userCol.UpdateOne(ctx, filter, updateData)

	return err
}