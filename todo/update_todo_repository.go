package todo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateTodoById(db *mongo.Database) func(context.Context, Todo) error {
	return func(ctx context.Context, todo Todo) error {
		collection := getTodosCollection(db)
		filter := bson.M{"_id": todo.ID}
		rs, err := collection.UpdateOne(ctx, filter, bson.M{"$set": todo})

		if rs.ModifiedCount == 0 {
			return errors.New("todo can not update")
		}
		return err
	}
}
