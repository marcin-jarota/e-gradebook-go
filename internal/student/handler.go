package student

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	transport.Handler
	studentService ports.StudentService
	markService    ports.MarkService
}

func NewStudentHandler(service ports.StudentService, markService ports.MarkService) *StudentHandler {
	return &StudentHandler{
		studentService: service,
		markService:    markService,
	}
}

func (h *StudentHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/students", auth.IsAuthenticatedByHeader())
	r.Get("/", auth.UserIs("teacher", "admin"), h.GetAll)
	r.Get("/:studentID/marks", auth.UserIs(domain.AdminRole, domain.StudentRole), h.GetMarks)
}

func (h *StudentHandler) GetAll(c *fiber.Ctx) error {
	students, err := h.studentService.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, students)
}

func (h *StudentHandler) GetMarks(c *fiber.Ctx) error {
	studentID, err := h.ParseIntParam(c.Params("studentId"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	marks, err := h.markService.GetByStudent(studentID)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, marks)
}

func (h *StudentHandler) GetAllStudents(c *fiber.Ctx) error {
	studentsList, err := h.studentService.GetAll()

	if err != nil {
		return h.JSONError(c, err, 500)
	}

	return h.JSON(c, studentsList)
}
