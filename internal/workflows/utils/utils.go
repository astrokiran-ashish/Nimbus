package workflowutils

import (
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
)

func GetStartWorkflowOptions(workflowName, taskQueue, workflowID string) client.StartWorkflowOptions {
	return client.StartWorkflowOptions{
		ID:                       workflowID,
		TaskQueue:                taskQueue,
		WorkflowExecutionTimeout: time.Hour * 24,
		WorkflowTaskTimeout:      time.Minute * 1,
		WorkflowRunTimeout:       time.Hour,
	}
}

func GetActivityOptions() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 1,
		StartToCloseTimeout:    time.Minute * 1,
		HeartbeatTimeout:       time.Second * 1,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 1,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute * 1,
		},
	}
}

func GetChildWorkflowOptions(ctx workflow.Context, workflowID string) (workflow.Context, workflow.ChildWorkflowOptions) {
	childWorkflowOption := workflow.ChildWorkflowOptions{
		WorkflowExecutionTimeout: time.Hour * 24,
		WorkflowTaskTimeout:      time.Minute * 1,
		WorkflowRunTimeout:       time.Hour,
		WorkflowID:               workflowID,
	}

	childCtx := workflow.WithChildOptions(ctx, childWorkflowOption)
	return childCtx, childWorkflowOption
}

func CalculateTalkTime(con workflowstates.ConsultationState) time.Duration {
	return time.Duration(10) * time.Minute
	// fractionalMins := (con.UserWalletBalance % con.ConsultantPricePerMinute) / con.ConsultantPricePerMinute
	// fullMins := (con.UserWalletBalance - fractionalMins) / con.ConsultantPricePerMinute
	// // Add 1 minute for the first minute of the call
	// return time.Duration(fullMins)*time.Minute + time.Duration(fractionalMins)*time.Second + time.Duration(1)*time.Minute
}
