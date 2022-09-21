package todo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deleteTodoByIdHandlerFn func(context.Context, primitive.ObjectID) error

func (fn deleteTodoByIdHandlerFn) DeleteTodoById(ctx context.Context, objectId primitive.ObjectID) error {
	return fn(ctx, objectId)
}

func DeleteTodoByIdHandler(svc deleteTodoByIdHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := c.Param("todo")
		objectId, err := primitive.ObjectIDFromHex(todo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = svc.DeleteTodoById(c.Request().Context(), objectId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, "delete todo success")
	}
}

// type deleteTodoByIdHandlerFn func(context.Context, string) error

// func (fn deleteTodoByIdHandlerFn) DeleteTodoById(ctx context.Context, str string) error {
// 	return fn(ctx, str)
// }

// func DeleteTodoByIdHandler(svc deleteTodoByIdHandlerFn) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		todo := c.Param("todo")
// 		if err := svc.DeleteTodoById(c.Request().Context(), todo); err != nil {
// 			return c.String(http.StatusBadRequest, err.Error())
// 		}
// 		return c.JSON(http.StatusOK, "deleted")
// 	}
// }
