package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection name
const (
	users = "users"
)

// UserCol ...
func UserCol() *mongo.Collection {
	return db.Collection(users)
}
