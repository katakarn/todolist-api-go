package todo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func getTodosCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("todos")
}
