package schoolyear

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type SchoolYearHandler struct {
	transport.Handler
	service ports.SchoolYearService
}

func NewSchoolYearHandler(service ports.SchoolYearService) *SchoolYearHandler {
	return &SchoolYearHandler{
		service: service,
	}
}

func (h *SchoolYearHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/school-year", auth.IsAuthenticatedByHeader())

	r.Get("/", auth.UserIs("admin", "teacher"), h.GetAll)
}

func (h *SchoolYearHandler) GetAll(c *fiber.Ctx) error {
	schoolYears, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, schoolYears)
}
