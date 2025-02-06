package workflowutils

import (
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
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

func GetChildWorkflowOptions(workflowID string) workflow.ChildWorkflowOptions {
	return workflow.ChildWorkflowOptions{
		WorkflowExecutionTimeout: time.Hour * 24,
		WorkflowTaskTimeout:      time.Minute * 1,
		WorkflowRunTimeout:       time.Hour,
		WorkflowID:               workflowID,
	}
}
