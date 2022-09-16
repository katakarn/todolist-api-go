package main

import (
	"net/http"

	"context"

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
		AuthSource:    "test",
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
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", getAllUser)
	e.GET("/users/:id", getUserById)
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getAllUser(c echo.Context) error {
	// Get user from database
	return c.String(http.StatusOK, "User")
}

// func GetUserByUsername(db *mongo.Database) func(context.Context, string) (*User, error) {
// 	return func(ctx context.Context, username string) (*User, error) {
// 		collection := getUsersCollection(db)
// 		filter := bson.M{"username": username}
// 		var user User
// 		if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
// 			return nil, err
// 		}
// 		return &user, nil
// 	}
// }

// e.GET("/users/:id", getUser)
func getUserById(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	// Save user to database
	return c.String(http.StatusCreated, "User created")
}

func updateUser(c echo.Context) error {
	// Update user in database
	return c.String(http.StatusOK, "User updated")
}

func deleteUser(c echo.Context) error {
	// Delete user from database
	return c.String(http.StatusOK, "User deleted")
}

// Run the server
// $ go run main.go

// Test the server
// $ curl -i http://localhost:1323
// $ curl -i http://localhost:1323/users/1
// $ curl -i -X POST http://localhost:1323/users
// $ curl -i -X PUT http://localhost:1323/users/1
// $ curl -i -X DELETE http://localhost:1323/users/1
