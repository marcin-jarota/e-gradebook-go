package notification

import (
	"e-student/internal/app/ports"
	"errors"
	"fmt"
	"time"
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

func (s *NotificationService) GetNotificationsForUser(userID int) ([]ports.Notification, error) {
	var output []ports.Notification
	list, err := s.repo.GetNotificationsForUser(userID)

	if err != nil {
		return output, errors.New("notificaiton.errorFetch.internal")
	}

	for _, n := range list {
		output = append(output, ports.Notification{
			Read:      n.Read,
			Message:   n.Message,
			CreatedAt: n.CreatedAt.Format(time.RFC3339),
		})
	}

	return output, nil
}
