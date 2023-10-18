package auth

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	transport.Handler
	service ports.AuthService
}

func NewAuthHandler(service ports.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) BindRouting(app *fiber.App, auth *middleware.AuthMiddleware) *AuthHandler {
	app.Post("/login", h.Login)
	app.Get("/logout", auth.IsAuthenticatedByHeader(), h.Logout)

	return h
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	payload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "Could not parse message",
		})
	}

	token, err := h.service.Login(payload.Email, payload.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.User)
	err := h.service.Logout(int(user.ID))
	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return c.SendStatus(200)

}
