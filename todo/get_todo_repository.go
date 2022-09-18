package todo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTodoById(db *mongo.Database) func(context.Context, string) (*Todo, error) {
	return func(ctx context.Context, str string) (*Todo, error) {
		collection := getTodosCollection(db)
		filter := bson.M{"title": str}
		var todo Todo
		if err := collection.FindOne(ctx, filter).Decode(&todo); err != nil {
			return nil, err
		}
		return &todo, nil
	}
}

func GetAllTodo(db *mongo.Database) func(context.Context) ([]*Todo, error) {
	return func(ctx context.Context) ([]*Todo, error) {
		collection := getTodosCollection(db)
		var todos []*Todo
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
		if err = cursor.All(ctx, &todos); err != nil {
			return nil, err
		}
		return todos, nil
	}
}
