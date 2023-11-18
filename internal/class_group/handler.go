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
	r := app.Group("/class", auth.IsAuthenticatedByHeader())
	r.Get("/all", auth.IsAdmin(), h.GetAll)
	r.Post("/create", auth.IsAdmin(), h.Create)
}

func (h *ClassGroupHandler) Create(c *fiber.Ctx) error {
	var p ports.AddClassGroupInput

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := h.service.AddClassGroup(&p); err != nil {

		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	return h.JSON(c, fiber.Map{
		"success": true,
	})
}

func (h *ClassGroupHandler) GetAll(c *fiber.Ctx) error {
	classGroups, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, classGroups)
}
