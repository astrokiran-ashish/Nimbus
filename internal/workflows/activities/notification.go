package activities

import (
	"context"

	"github.com/google/uuid"
)

func (a *Activities) SendNotificationToConsultant(ctx context.Context, consultantID uuid.UUID, message string) error {
	if err := a.notification.SendNotificationToConsultant(consultantID, message); err != nil {
		return err
	}
	return nil
}

func (a *Activities) SendNotificationToUser(ctx context.Context, userID uuid.UUID, message string) error {
	if err := a.notification.SendNotificationToUser(userID, message); err != nil {
		return err
	}
	return nil
}
