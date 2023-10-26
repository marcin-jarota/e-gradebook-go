package user

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	transport.Handler
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) BindRouting(app *fiber.App, auth *middleware.AuthMiddleware) {
	r := app.Group("/user", auth.IsAuthenticatedByHeader(), auth.IsAdmin())

	r.Get("/list", h.GetAll)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, users)
}
