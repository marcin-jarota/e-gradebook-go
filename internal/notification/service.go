package notification

import (
	"e-student/internal/app/ports"
	"errors"
	"fmt"
)

type NotificationService struct {
	repo       ports.NotificationRepository
	strategies map[string]ports.NotificationStrategy
}

func NewNotificationService(strategies map[string]ports.NotificationStrategy, repo ports.NotificationRepository) *NotificationService {
	return &NotificationService{
		strategies: strategies,
		repo:       repo,
	}
}

func (s *NotificationService) SendNotification(notificationType string, notification ports.Notification) error {
	stratery, ok := s.strategies[notificationType]
	if !ok {
		return errors.New(fmt.Sprintf("Missing implementation for required type: %s", notificationType))
	}

	return stratery.SendNotification(notification)
}
