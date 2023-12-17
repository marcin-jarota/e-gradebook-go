package ports

import "e-student/internal/app/domain"

type (
	NotificationRepository interface {
		CreateNotification(notification domain.Notification) error
		GetNotificationsForUser(userID int) ([]domain.Notification, error)
	}

	Notification struct {
		UserID    int    `json:"userID,omitempty"`
		Message   string `json:"message"`
		Read      bool   `json:"read"`
		CreatedAt string `json:"createdAt,omitempty"`
	}

	NotificationService interface {
		SendNotification(notificationType string, notification Notification) error
		GetNotificationsForUser(userID int) ([]Notification, error)
	}

	NotificationStrategy interface {
		SendNotification(n Notification) error
	}
)
