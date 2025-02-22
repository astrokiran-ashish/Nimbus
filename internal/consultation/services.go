package consultation

import (
	"context"
	"fmt"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	worflowutils "github.com/astrokiran/nimbus/internal/workflows/utils"
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"

	"github.com/google/uuid"
)

func (con *Consultation) HandleConsultationCreation(req CreateConsultationRequest) (*model.Consultation, error) {
	consultation, err := con.CreateConsultation(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create consultation: %w", err)
	}

	consultationState := toConsultationState(consultation)
	workflowID := fmt.Sprintf("%s:%s", constants.ConsultationLifecycleWorkflowName, consultation.ConsultationID.String())
	worflowOptions := worflowutils.GetStartWorkflowOptions(constants.ConsultationLifecycleWorkflowName, con.workflowTaskQueue, workflowID)

	_, err = con.engine.Client.ExecuteWorkflow(context.Background(), worflowOptions, constants.ConsultationLifecycleWorkflowName, consultationState)
	if err != nil {
		return nil, fmt.Errorf("failed to start workflow: %w", err)
	}

	return consultation, nil
}

func (con *Consultation) HandleGetConsultation(consultationID uuid.UUID) (*model.Consultation, error) {
	consultation, err := con.GetConsultation(consultationID)
	if err != nil {
		return nil, fmt.Errorf("failed to get consultation: %w", err)
	}
	return consultation, nil
}

func (con *Consultation) HandleUpdateConsultation(req UpdateConsultationRequest) (*model.Consultation, error) {
	consultation, err := con.UpdateConsultation(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update consultation: %w", err)
	}
	return consultation, nil
}

func (con *Consultation) HandleConsultantActionEvent(req ConsultantActionEventRequest) error {
	workflowID, signalCh, err := EventRouter(req)
	if err != nil {
		fmt.Println("failed to route event: %w", err)
		return fmt.Errorf("failed to route event: %w", err)
	}
	consultantActionEvent := workflowstates.ConsultantActionEvent{
		ConsultationID: req.ConsultationID,
		Action:         req.Action,
	}
	cxt := context.Background()
	err = con.engine.Client.SignalWorkflow(cxt, workflowID, "", signalCh, consultantActionEvent)
	if err != nil {
		return fmt.Errorf("failed to signal workflow: %w", err)
	}
	return nil
}
