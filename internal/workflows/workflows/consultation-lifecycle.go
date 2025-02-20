package workflows

import (
	"fmt"

	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	utils "github.com/astrokiran/nimbus/internal/workflows/utils"
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
	"go.temporal.io/sdk/workflow"
)

func ConsultationLifecycleWorkflow(ctx workflow.Context, state workflowstates.ConsultationState) error {
	workflow.GetLogger(ctx).Info("ConsultationWorkflow started", "state", state)

	workflowID := fmt.Sprintf("%s:%s", constants.ConsultationHandShakeWorkflowName, state.ConsultationID)
	fmt.Println("workflowID", workflowID)
	childCtx, _ := utils.GetChildWorkflowOptions(ctx, workflowID)
	err := workflow.ExecuteChildWorkflow(childCtx, constants.ConsultationHandShakeWorkflowName, state).Get(childCtx, nil)
	if err != nil {
		return err
	}

	workflowID = fmt.Sprintf("%s:%s", constants.ConsultationSessionWorkflowName, state.ConsultationID)
	childCtx, _ = utils.GetChildWorkflowOptions(ctx, workflowID)
	err = workflow.ExecuteChildWorkflow(childCtx, constants.ConsultationSessionWorkflowName, state).Get(childCtx, nil)
	if err != nil {
		return err
	}

	workflowID = fmt.Sprintf("%s:%s", constants.ConsultationBillingWorkflowName, state.ConsultationID)
	childCtx, _ = utils.GetChildWorkflowOptions(ctx, workflowID)
	err = workflow.ExecuteChildWorkflow(childCtx, constants.ConsultationBillingWorkflowName, state).Get(childCtx, nil)
	if err != nil {
		return err
	}

	return nil
}
