package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huuloc2026/go-social/internal/core/ports"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(app *fiber.App, service ports.UserService) {
	handler := &UserHandler{service: service}
	app.Post("/users", handler.CreateUser)
	app.Get("/users/:id", handler.GetUser)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	req := new(request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	user, err := h.service.CreateUser(req.Username, req.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}
