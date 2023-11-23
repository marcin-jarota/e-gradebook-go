package mark

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type markHandler struct {
	transport.Handler
	markService ports.MarkService
}

func NewMarkHandler(markService ports.MarkService) *markHandler {
	return &markHandler{
		markService: markService,
	}
}

func (h *markHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/marks", auth.IsAuthenticatedByHeader())

	r.Post("/", h.CreateMark)
}

func (h *markHandler) CreateMark(c *fiber.Ctx) error {
	var p ports.MarkCreatePayload

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := h.markService.CreateMark(p); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, fiber.Map{
		"success": true,
	})
}
