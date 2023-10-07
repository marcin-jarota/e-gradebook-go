package http

import (
	"e-student/internal/app/ports"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	service ports.StudentService
}

func NewStudentHandler(srv ports.StudentService) *StudentHandler {
	return &StudentHandler{
		service: srv,
	}
}

func (s *StudentHandler) GetAllStudents(c *fiber.Ctx) error {
	studentsList, err := s.service.GetAll()

	if err != nil {
		return err
	}

	return c.Render("pages/students", fiber.Map{"Students": studentsList}, "layouts/main")
}
