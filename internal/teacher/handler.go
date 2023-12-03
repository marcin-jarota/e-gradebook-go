package teacher

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type TeacherHandler struct {
	transport.Handler
	teacherService ports.TeacherService
}

func NewTeacherHandler(teacherSrvc ports.TeacherService) *TeacherHandler {
	return &TeacherHandler{
		teacherService: teacherSrvc,
	}
}

func (h *TeacherHandler) BindRouting(app *fiber.App, auth *middleware.AuthMiddleware) {
	r := app.Group("/teachers")

	r.Get("/", h.List)
}

func (h *TeacherHandler) List(c *fiber.Ctx) error {
	output, err := h.teacherService.GetAll()

	if err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	return h.JSON(c, output)
}
