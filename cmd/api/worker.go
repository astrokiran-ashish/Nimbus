package main

import (
	"fmt"
	"log"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/astrokiran/nimbus/internal/notification"
	"github.com/astrokiran/nimbus/internal/workflows/activities"
	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	workflows "github.com/astrokiran/nimbus/internal/workflows/workflows"
)

func (app *application) startWorker(c client.Client, taskQueue string, notification notification.INotification) error {
	// Validate client connection
	if c == nil {
		return fmt.Errorf("temporal client is nil")
	}

	log.Printf("Starting worker for task queue: %s", taskQueue)
	// Create worker instance
	w := worker.New(c, taskQueue, worker.Options{})
	// Register workflow
	w.RegisterWorkflowWithOptions(workflows.ConsultationWorkflow,
		workflow.RegisterOptions{Name: constants.ConsultationWorkflowName})
	w.RegisterWorkflowWithOptions(workflows.ConsultantStartWorkflow,
		workflow.RegisterOptions{Name: constants.ConsultantStartWorkflowName})

	activityInstance := activities.NewActivities(notification)
	w.RegisterActivityWithOptions(activityInstance.SendNotificationToConsultant, activity.RegisterOptions{
		Name: constants.SendNotificationToConsultantActivity,
	})
	// Run worker in a separate goroutine to avoid blocking
	go func() {
		if err := w.Run(worker.InterruptCh()); err != nil {
			log.Fatalf("Failed to start Temporal worker: %v", err)
		}
	}()

	return nil
}

func CreateTemporalClient(temporalHostPort string) (client.Client, error) {

	// Create a Temporal client
	c, err := client.Dial(client.Options{
		HostPort: temporalHostPort,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("Connected to Temporal at %s\n", temporalHostPort)

	return c, nil
}
