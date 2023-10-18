package middleware

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	transport.Handler
	auth ports.AuthService
}

func NewAuthMiddleware(auth ports.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		auth: auth,
	}

}

func (m *AuthMiddleware) IsAuthenticatedByHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		parts := strings.Split(authHeader, " ")

		if len(parts) < 2 {
			return m.JSONError(c, errors.New("missing_token"), fiber.StatusBadRequest)
		}

		schema := parts[0]
		token := parts[1]

		if schema != "Bearer" || token == "" {
			return m.JSONError(c, errors.New("unauthorized"), fiber.StatusUnauthorized)
		}

		isLoggedIn, user := m.auth.IsLoggedIn(token)

		if isLoggedIn {
			c.Locals("user", user)
			return c.Next()
		}

		return m.JSONError(c, errors.New("unauthorized"), fiber.StatusUnauthorized)
	}
}

func (m *AuthMiddleware) IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, ok := c.Locals("user").(*domain.User)

		if !ok {
			return m.JSONError(c, errors.New("no session user"), fiber.StatusUnauthorized)
		}

		if !u.IsAdmin() {
			return m.JSONError(c, errors.New("unauthorized"), fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}

func (m *AuthMiddleware) IsStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, ok := c.Locals("user").(*domain.User)

		if !ok {
			return m.JSONError(c, errors.New("no session user"), fiber.StatusUnauthorized)
		}

		if !u.IsStudent() {
			return m.JSONError(c, errors.New("unauthorized"), fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
