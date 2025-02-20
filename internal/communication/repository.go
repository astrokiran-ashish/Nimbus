package communication

import (
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
)

func (com *Communication) SaveEvent(event *model.AgoraWebhookEvents) error {
	// Save event to database using go-jet
	stmt := table.AgoraWebhookEvents.INSERT(
		table.AgoraWebhookEvents.AppID,
		table.AgoraWebhookEvents.ChannelName,
		table.AgoraWebhookEvents.EventTime,
		table.AgoraWebhookEvents.EventType,
		table.AgoraWebhookEvents.UID,
		table.AgoraWebhookEvents.CreatedAt,
		table.AgoraWebhookEvents.UpdatedAt,
		table.AgoraWebhookEvents.SessionID,
		table.AgoraWebhookEvents.Payload,
	).VALUES(
		event.AppID,
		event.ChannelName,
		event.EventTime,
		event.EventType,
		event.UID,
		event.CreatedAt,
		event.UpdatedAt,
		event.SessionID,
		event.Payload,
	)

	_, err := stmt.Exec(com.db.Conn)
	if err != nil {
		return err
	}
	return nil

}
