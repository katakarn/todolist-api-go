package todo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllTodo(db *mongo.Database) func(context.Context, string) (*Todo, error) {
	return func(ctx context.Context, str string) (*Todo, error) {
		collection := getUsersCollection(db)
		filter := bson.M{"title": str}
		var todo Todo
		if err := collection.FindOne(ctx, filter).Decode(&todo); err != nil {
			return nil, err
		}
		return &todo, nil
	}
}
