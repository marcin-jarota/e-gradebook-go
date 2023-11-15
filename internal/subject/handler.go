package subject

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"
	"errors"
	"strconv"

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
	r := app.Group("/subject", auth.IsAuthenticatedByHeader())

	r.Get("/all", auth.IsAdmin(), h.GetAll)
	r.Get("/delete/:id", auth.IsAdmin(), h.DeleteSubject)
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

func (h *SubjectHandler) DeleteSubject(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil {
		return h.JSONError(c, errors.New("subject.error.id"), fiber.StatusBadRequest)
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, nil)
}

func (h *SubjectHandler) GetAll(c *fiber.Ctx) error {
	response, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, response)
}
