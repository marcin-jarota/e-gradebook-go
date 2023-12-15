package notification

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
)

type InAppNotification struct {
	repo ports.NotificationRepository
}

func NewInAppNotification(repo ports.NotificationRepository) *InAppNotification {
	return &InAppNotification{repo: repo}
}

func (s *InAppNotification) SendNotification(notification ports.Notification) error {
	return s.repo.CreateNotification(domain.Notification{
		UserID:  uint(notification.UserID),
		Message: notification.Message,
		Read:    false,
	})
}
