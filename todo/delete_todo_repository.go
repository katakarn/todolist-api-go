package todo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteTodoById(db *mongo.Database) func(context.Context, primitive.ObjectID) error {
	return func(ctx context.Context, objectId primitive.ObjectID) error {
		collection := getTodosCollection(db)
		filter := bson.M{"_id": objectId}
		rs, err := collection.DeleteOne(ctx, filter)
		if rs.DeletedCount == 0 {
			return errors.New("todo can not delete")
		}
		return err
	}
}
