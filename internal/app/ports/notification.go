package ports

import "e-student/internal/app/domain"

type (
	NotificationRepository interface {
		CreateNotification(notification domain.Notification) error
		GetNotificationsForUser(userID int) ([]domain.Notification, error)
	}

	Notification struct {
		UserID  int    `json:"userID"`
		Message string `json:"message"`
		Read    bool   `json:"read"`
	}

	NotificationService interface {
		SendNotification(notificationType string, notification Notification) error
	}

	NotificationStrategy interface {
		SendNotification(n Notification) error
	}
)
