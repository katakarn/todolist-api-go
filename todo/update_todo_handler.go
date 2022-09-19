package todo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type updateTodoHandlerFn func(context.Context, Todo) error

func (fn updateTodoHandlerFn) UpdateTodoById(ctx context.Context, todo Todo) error {
	return fn(ctx, todo)
}

func UpdateTodoHandler(svc updateTodoHandlerFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo Todo
		if err := c.Bind(&todo); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := svc.UpdateTodoById(c.Request().Context(), todo); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, todo)
	}
}

// func UpdateTodoHandler(svc updateTodoHandlerFn) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		todoId := c.Param("todo")

// 		err := svc.UpdateTodoById(c.Request().Context(), todoId, &Todo{})
// 		if err != nil {
// 			return c.String(http.StatusBadRequest, err.Error())
// 		}
// 		return c.JSON(http.StatusOK, "updated")
// 	}
// }
