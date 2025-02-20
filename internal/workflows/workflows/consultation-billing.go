package workflows

import (
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
	"go.temporal.io/sdk/workflow"
)

func ConsultationBillingWorkflow(ctx workflow.Context, state workflowstates.ConsultationState) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("ConsultationBillingWorkflow started", "state", state)
	defer logger.Info("ConsultationBillingWorkflow completed", "state", state)

	return nil
}
