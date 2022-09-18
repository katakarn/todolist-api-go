package todo

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getTodoByIdFn func(context.Context, string) (*Todo, error)

func (gn getTodoByIdFn) GetTodoById(ctx context.Context, str string) (*Todo, error) {
	return gn(ctx, str)
}

func GetTodoByIdHandler(svc getTodoByIdFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := c.Param("todo")
		todos, err := svc.GetTodoById(c.Request().Context(), todo)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, todos)
	}
}

type getAllTodoFn func(context.Context) ([]*Todo, error)

func (gn getAllTodoFn) GetAllTodo(ctx context.Context) ([]*Todo, error) {
	return gn(ctx)
}

func GetAllTodoHandler(svc getAllTodoFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := svc.GetAllTodo(c.Request().Context())
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, todos)
	}
}

