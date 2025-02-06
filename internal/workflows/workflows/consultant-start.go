package workflows

import (
	"go.temporal.io/sdk/workflow"

	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	worflowutils "github.com/astrokiran/nimbus/internal/workflows/utils"
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
)

func ConsultantStartWorkflow(ctx workflow.Context, state workflowstates.ConsultationState) error {
	workflow.GetLogger(ctx).Info("ConsultantStartWorkflow started", "state", state)
	activityOptions := worflowutils.GetActivityOptions()
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	err := workflow.ExecuteActivity(ctx, constants.SendNotificationToConsultantActivity, state.ConsultantID, "Consultation started").Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
