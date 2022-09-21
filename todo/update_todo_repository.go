package todo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateTodoByIdField(db *mongo.Database) func(context.Context, Todo) error {
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

func UpdateTodoByIdParam(db *mongo.Database) func(context.Context, primitive.ObjectID, Todo) error {
	return func(ctx context.Context, objectId primitive.ObjectID, todo Todo) error {
		collection := getTodosCollection(db)
		filter := bson.M{"_id": objectId}
		rs, err := collection.UpdateOne(ctx, filter, bson.M{"$set": todo})
		// rs, err := collection.UpdateOne(ctx, filter, bson.M{"title": todo.Title, "description": todo.Description})

		if rs.ModifiedCount == 0 {
			return errors.New("todo can not update")
		}
		return err
	}
}
