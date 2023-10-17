package http

import (
	"e-student/internal/app/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	Handler
	service ports.StudentService
}

func NewStudentHandler(srv ports.StudentService) *StudentHandler {
	return &StudentHandler{
		service: srv,
	}
}

func (h *StudentHandler) BindRouting(app fiber.Router) {
	app.Get("/student/marks/:studentId", h.GetMarks)
	app.Get("/students", h.GetAllStudents)
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
