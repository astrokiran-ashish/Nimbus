package communication

import (
	"context"
	"fmt"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	"go.uber.org/zap"
)

func (com *Communication) HandleWebhookEvent(event *model.AgoraWebhookEvents) error {
	com.logger.Info("Received webhook event", zap.Any("event", event))
	if err := com.SaveEvent(event); err != nil {
		return err
	}
	workflowID := fmt.Sprintf("%s:%s", constants.AgoraWehbookEventSignalCh, *event.SessionID)
	cxt := context.Background()
	err := com.engine.Client.SignalWorkflow(cxt, workflowID, "", constants.SessionEndSignalCh, true)
	if err != nil {
		return fmt.Errorf("failed to signal workflow: %w", err)
	}
	return nil
}
