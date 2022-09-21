package todo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type updateTodoByIdFieldHandlerFn func(context.Context, Todo) error

func (fn updateTodoByIdFieldHandlerFn) UpdateTodoByIdField(ctx context.Context, todo Todo) error {
	return fn(ctx, todo)
}

func UpdateTodoByIdFieldHandler(svc updateTodoByIdFieldHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := Todo{}
		if err := c.Bind(&todo); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := svc.UpdateTodoByIdField(c.Request().Context(), todo); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, todo)
	}
}

type updateTodoByIdParamHandlerFn func(context.Context, primitive.ObjectID, Todo) error

func (fn updateTodoByIdParamHandlerFn) UpdateTodoByIdParam(ctx context.Context, objectId primitive.ObjectID, todo Todo) error {
	return fn(ctx, objectId, todo)
}

func UpdateTodoByIdParamHandler(svc updateTodoByIdParamHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := Todo{}
		todoId := c.Param("todo")
		objectId, err := primitive.ObjectIDFromHex(todoId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		if err := c.Bind(&todo); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		err = svc.UpdateTodoByIdParam(c.Request().Context(), objectId, todo)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, todo)
	}
}
