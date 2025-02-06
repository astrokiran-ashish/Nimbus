package activities

import "github.com/astrokiran/nimbus/internal/notification"

type Activities struct {
	notification notification.INotification
}

func NewActivities(notification notification.INotification) *Activities {
	return &Activities{notification: notification}
}
