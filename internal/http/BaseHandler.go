package http

import (
	"e-student/internal/app/domain"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func (h *Handler) JSONError(c *fiber.Ctx, err error, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func (h *Handler) JSON(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"data": data,
	})
}

func (h *Handler) RenderWithGlobals(c *fiber.Ctx, tplName string, data interface{}, layouts ...string) error {
	user := c.Locals("user").(*domain.User)

	globals := map[string]interface{}{
		"User": user,
	}

	parsed, ok := data.(fiber.Map)
	if ok {
		for k, v := range parsed {
			globals[k] = v
		}

	}

	return c.Render(tplName, globals, layouts...)
}
