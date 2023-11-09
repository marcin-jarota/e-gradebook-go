package auth

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"
	"errors"
	"strings"

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
	app.Get("/token-valid", h.IsTokenValid)
	app.Get("/logout", auth.IsAuthenticatedByHeader(), h.Logout)

	return h
}

func (h *AuthHandler) IsTokenValid(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	parts := strings.Split(authHeader, " ")

	if len(parts) < 2 {
		return h.JSONError(c, errors.New("missing_token"), fiber.StatusBadRequest)
	}

	schema := parts[0]
	token := parts[1]

	if schema != "Bearer" || token == "" {
		return h.JSONError(c, errors.New("unauthorized"), fiber.StatusUnauthorized)
	}

	isValid, _ := h.service.IsTokenValid(token)

	if !isValid {
		return h.JSONError(c, errors.New("invalid token"), fiber.StatusUnauthorized)
	}

	return h.JSON(c, nil)
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
