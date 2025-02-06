package workflows

import (
	"fmt"

	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	utils "github.com/astrokiran/nimbus/internal/workflows/utils"
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
	"go.temporal.io/sdk/workflow"
)

func ConsultationWorkflow(ctx workflow.Context, state workflowstates.ConsultationState) error {
	workflow.GetLogger(ctx).Info("ConsultationWorkflow started", "state", state)

	workflowID := fmt.Sprintf("%s:%s", constants.ConsultantStartWorkflowName, state.ConsultationID)
	ChildWorkflowOptions := utils.GetChildWorkflowOptions(workflowID)
	err := workflow.ExecuteChildWorkflow(ctx, constants.ConsultantStartWorkflowName, ChildWorkflowOptions, state).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
