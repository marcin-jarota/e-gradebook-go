package classgroup

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type ClassGroupHandler struct {
	transport.Handler
	service ports.ClassGroupService
}

func NewClassGroupHandler(service ports.ClassGroupService) *ClassGroupHandler {
	return &ClassGroupHandler{
		service: service,
	}
}

func (h *ClassGroupHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/class")
	r.Get("/all", h.GetAll)
}

func (h *ClassGroupHandler) GetAll(c *fiber.Ctx) error {
	classGroups, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, classGroups)
}
