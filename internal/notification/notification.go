package notification

import (
	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type INotification interface {
	SendNotificationToConsultant(consultantID uuid.UUID, message string) error
	SendNotificationToUser(userID uuid.UUID, message string) error
}

type Notification struct {
	db     *database.Database
	logger *zap.Logger
}

func NewNotification(db *database.Database, logger *zap.Logger) *Notification {
	return &Notification{db: db, logger: logger}
}

func (n *Notification) SendNotificationToConsultant(consultantID uuid.UUID, message string) error {
	n.logger.Info("Sending notification to consultant", zap.String("message", message))
	return nil
}

func (n *Notification) SendNotificationToUser(userID uuid.UUID, message string) error {
	n.logger.Info("Sending notification to user", zap.String("message", message))
	return nil
}
