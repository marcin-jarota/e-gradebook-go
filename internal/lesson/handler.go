package lesson

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type LessonHandler struct {
	transport.Handler
	lessonService ports.LessonService
}

func NewLessonHandler(lessonService ports.LessonService) *LessonHandler {
	return &LessonHandler{
		lessonService: lessonService,
	}
}

func (h *LessonHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/lessons", auth.IsAuthenticatedByHeader())

	r.Post("/", auth.IsAdmin(), h.Create)
}

func (h *LessonHandler) Create(c *fiber.Ctx) error {
	var p ports.CreateLessonPayload

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := h.lessonService.Create(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	return h.JSON(c, fiber.Map{"ok": true})
}
