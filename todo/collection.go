package todo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func getUsersCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("todos")
}
