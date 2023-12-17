package notification

import (
	"e-student/internal/adapters/transport"
	"e-student/internal/app/ports"
	"e-student/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type NotificationHandler struct {
	transport.Handler
	notificationService ports.NotificationService
}

func NewNotificationHandler(notificationService ports.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

func (h *NotificationHandler) BindRouting(app fiber.Router, auth *middleware.AuthMiddleware) {
	r := app.Group("/notifications", auth.IsAuthenticatedByHeader())

	r.Get("/:id", h.GetByUser)
}

func (h *NotificationHandler) GetByUser(c *fiber.Ctx) error {
	userID, err := h.ParseIntParam(c.Params("id", "0"))

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	notifications, err := h.notificationService.GetNotificationsForUser(userID)

	if err != nil {
		return h.JSONError(c, err, fiber.StatusBadRequest)
	}

	return h.JSON(c, notifications)
}
