package http

import (
	"e-student/internal/app/ports"

	"github.com/gofiber/fiber/v2"
)

type MarksHandler struct {
	marksRepository ports.MarkRepository
}

func (h *MarksHandler) BindRouting(app *fiber.App) {
	// app.Get("/marks/:studentId", h.GetMarksByStudent)
}

// func (h *MarksHandler) GetMarksByStudent(c *fiber.Ctx) error {

// }
