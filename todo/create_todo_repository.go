package todo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTodo(db *mongo.Database) func(context.Context, *Todo) error {
	return func(ctx context.Context, todo *Todo) error {
		collection := getTodosCollection(db)
		_, err := collection.InsertOne(ctx, todo)
		if err != nil {
			return err
		}
		return nil
	}
}
