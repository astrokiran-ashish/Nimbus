package workflows

import (
	"time"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	utils "github.com/astrokiran/nimbus/internal/workflows/utils"
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
	"go.temporal.io/sdk/workflow"
)

func ConsultationSessionWorkflow(ctx workflow.Context, state workflowstates.ConsultationState) (time.Duration, error) {
	// Define the workflow

	logger := workflow.GetLogger(ctx)
	logger.Info("ConsultationSessionWorkflow started", "state", state)
	defer logger.Info("ConsultationSessionWorkflow completed", "state", state)

	// fetch wallet balance of user
	talkTimeDuration := utils.CalculateTalkTime(state)
	logger.Info("Talk time calculated", "talkTime", talkTimeDuration)

	start := time.Now()

	talkTime := workflow.NewTimer(ctx, talkTimeDuration)
	selector := workflow.NewSelector(ctx)
	selector.AddFuture(talkTime, func(f workflow.Future) {
		logger.Info("Talk time finished")
	})

	wehbookEventChannel := workflow.GetSignalChannel(ctx, constants.SessionEndSignalCh)
	var wehbookEvent *model.AgoraWebhookEvents
	selector.AddReceive(wehbookEventChannel, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &wehbookEvent)
		logger.Info("Recevied webhook response", "response", wehbookEvent)
		// call activity to handle webhook event response
	})

	selector.Select(ctx)

	consultationTime := time.Since(start)

	return consultationTime, nil
}
