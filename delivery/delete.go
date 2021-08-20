package delivery

import (
	"todo/domain"

	"github.com/gofiber/fiber/v2"
)

type todoDeleteHandler struct {
	todoUsecase domain.TodoUsecase
}

func NewTodoDeleteHandler(c *fiber.App, th domain.TodoUsecase) {
	handler := &todoDeleteHandler{
		todoUsecase: th,
	}

	c.Post("/todo/delete", handler.Delete)
}

func (h *todoDeleteHandler) Delete(c *fiber.Ctx) error {
	todo := new(domain.Todo)
	err := c.BodyParser(todo)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Unexpected Request. To check your Todo data",
		})
	}

	err = h.todoUsecase.Delete(todo.ID)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON("Delete OK")
}
