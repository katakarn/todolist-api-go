package main

import (
	"net/http"

	"context"
	"time"

	"github.com/katakarn/todolist-api-go/todo"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	Client *mongo.Client
}

func InitMongoDB(ctx context.Context) *DB {
	// To configure auth via URI instead of a Credential, use
	// "mongodb://root:password@localhost:27017".
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "todolist",
		Username:      "root",
		Password:      "password",
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	conn := &DB{
		Client: client,
	}

	return conn
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := InitMongoDB(ctx)
	defer func() {
		if err := db.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongodb := db.Client.Database("todolist")

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/todos/:todo", todo.GetTodoByIdHandler(todo.GetTodoById(mongodb)))
	e.GET("/todos", todo.GetAllTodoHandler(todo.GetAllTodo(mongodb)))

	e.Logger.Fatal(e.Start(":1323"))
}
