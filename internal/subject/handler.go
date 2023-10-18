package subject

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type SubjectHandler struct {
	transport.Handler
	service ports.SubjectService
}

func NewSubjectHandler(service ports.SubjectService) *SubjectHandler {
	return &SubjectHandler{
		service: service,
	}
}

func (h *SubjectHandler) BindRouting(app *fiber.App, auth *middleware.AuthMiddleware) {
	r := app.Group("/subject", auth.IsAuthenticatedByHeader(), auth.IsAdmin())

	r.Get("/all", h.GetAll)
	r.Post("/create", h.AddSubject)
}

func (h *SubjectHandler) AddSubject(c *fiber.Ctx) error {
	var p ports.SubjectAddPayload

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	err := h.service.AddSubject(p.Name)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	return h.JSON(c, fiber.Map{"success": true})
}

func (h *SubjectHandler) GetAll(c *fiber.Ctx) error {
	response, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, response)
}
