package middleware

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
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
		fmt.Println(authHeader)
		parts := strings.Split(authHeader, " ")

		if len(parts) < 2 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		schema := parts[0]
		token := parts[1]

		fmt.Println(token)
		if schema != "Bearer" || token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		isLoggedIn, user := m.auth.IsLoggedIn(token)

		if isLoggedIn {
			c.Locals("user", user)
			return c.Next()
		}

		return c.SendStatus(fiber.StatusUnauthorized)
	}
}

func (m *AuthMiddleware) IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, ok := c.Locals("user").(*domain.User)

		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "no session user"})
		}

		if !u.IsAdmin() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"err": "user is not admin"})
		}

		return c.Next()
	}
}

func (m *AuthMiddleware) IsStudent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, ok := c.Locals("user").(*domain.User)

		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "no session user"})
		}

		if !u.IsStudent() {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		return c.Next()
	}
}
