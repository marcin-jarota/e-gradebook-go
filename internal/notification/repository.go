package notification

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormNotificationRepository struct {
	db *gorm.DB
}

func NewGormNotificationRepository(db *gorm.DB) *GormNotificationRepository {
	return &GormNotificationRepository{
		db: db,
	}
}

func (r *GormNotificationRepository) CreateNotification(notification domain.Notification) error {
	return r.db.Create(&notification).Error
}

func (r *GormNotificationRepository) GetNotificationsForUser(userID int) ([]domain.Notification, error) {
	var notifications []domain.Notification

	err := r.db.Find(&notifications, "user_id = ?", userID).Error

	if err != nil {
		return notifications, err
	}

	return notifications, nil
}
