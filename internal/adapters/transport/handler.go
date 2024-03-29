package transport

import (
	"e-student/internal/app/domain"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func (h *Handler) JSONError(c *fiber.Ctx, err error, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"data":  nil,
		"error": err.Error(),
	})
}

func (h *Handler) JSON(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"data":  data,
		"error": nil,
	})
}

func (h *Handler) ParseIntParam(param string) (int, error) {
	val, err := strconv.Atoi(param)

	if err != nil {
		return 0, errors.New("invalid.intParam")
	}

	return val, nil
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
