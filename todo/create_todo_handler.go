package todo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createTodoFn func(context.Context, *Todo) error

func (gn createTodoFn) CreateTodo(ctx context.Context, todo *Todo) error {
	return gn(ctx, todo)
}

func CreateTodoHandler(svc createTodoFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(Todo)
		if err := c.Bind(todo); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err := svc.CreateTodo(c.Request().Context(), todo); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, "todo created")
	}
}
