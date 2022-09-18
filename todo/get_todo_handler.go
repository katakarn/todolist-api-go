package todo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getTodoFn func(context.Context, string) (*Todo, error)

func (gn getTodoFn) GetAllTodo(ctx context.Context, str string) (*Todo, error) {
	return gn(ctx, str)
}

func GetTodoHandler(svc getTodoFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := c.Param("todo")
		todos, err := svc.GetAllTodo(c.Request().Context(), todo)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, todos)
	}
}
