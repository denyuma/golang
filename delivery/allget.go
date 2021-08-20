package delivery

import (
	"todo/domain"

	"github.com/gofiber/fiber/v2"
)

type todoAllGetHandler struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoAllGetHandler(c *fiber.App, th domain.TodoUsecase) {
	handler := &todoAllGetHandler{
		todoUsecase: th,
	}
	c.Get("/todos", handler.AllGet)
}

func (h *todoAllGetHandler) AllGet(c *fiber.Ctx) error {
	todos, err := h.todoUsecase.AllGet()
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(todos)
}
