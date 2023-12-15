package mark

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type markHandler struct {
	transport.Handler
	markService         ports.MarkService
	studnetService      ports.StudentService
	subjectService      ports.SubjectService
	notificationService ports.NotificationService
}

func NewMarkHandler(markService ports.MarkService, studnetService ports.StudentService, subjectService ports.SubjectService, notificationService ports.NotificationService) *markHandler {
	return &markHandler{
		markService:         markService,
		notificationService: notificationService,
		studnetService:      studnetService,
		subjectService:      subjectService,
	}
}

func (h *markHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/marks", auth.IsAuthenticatedByHeader())

	r.Post("/", auth.UserIs("teacher"), h.CreateMark)
}

func (h *markHandler) CreateMark(c *fiber.Ctx) error {
	var p ports.MarkCreatePayload

	if err := c.BodyParser(&p); err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	if err := h.markService.CreateMark(p); err != nil {
		return h.JSONError(c, err, fiber.StatusInternalServerError)
	}

	student, err := h.studnetService.GetOneByID(p.StudentID)

	if err != nil {
		log.Println("Could not read student", err)
	}

	subject, err := h.subjectService.GetOneByID(p.StudentID)
	if err != nil {
		log.Println("Could not read subject", err)
	}

	if student != nil && subject != nil {
		notification := ports.Notification{
			UserID:  student.UserID,
			Message: fmt.Sprintf("Nowa ocena z przedmiotu: %s, %0.2f", subject.Name, p.Value),
		}

		if err := h.notificationService.SendNotification(
			"in-app",
			notification,
		); err != nil {
			log.Println("Could not create notification ", err)
		}

	}

	return h.JSON(c, fiber.Map{
		"success": true,
	})
}
