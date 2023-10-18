package student

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	transport.Handler
	service    ports.StudentService
	middleware *middleware.AuthMiddleware
}

func NewStudentHandler(service ports.StudentService, middleware *middleware.AuthMiddleware) *StudentHandler {
	return &StudentHandler{
		service:    service,
		middleware: middleware,
	}
}

func (h *StudentHandler) BindRouting(app fiber.Router) {
	r := app.Group("/", h.middleware.IsAuthenticatedByHeader())

	r.Get("/:studentId/marks", h.middleware.IsStudent(), h.GetMarks)
}

func (h *StudentHandler) GetMarks(c *fiber.Ctx) error {
	studentId, err := strconv.Atoi(c.Params("studentId"))

	if err != nil {
		return h.JSONError(c, err, 500)
	}

	marks, err := h.service.GetMarks(studentId)

	if err != nil {
		return h.JSONError(c, err, 500)
	}

	return h.JSON(c, marks)
}

func (h *StudentHandler) GetAllStudents(c *fiber.Ctx) error {
	studentsList, err := h.service.GetAll()

	if err != nil {
		return h.JSONError(c, err, 500)
	}

	return h.JSON(c, studentsList)
}
