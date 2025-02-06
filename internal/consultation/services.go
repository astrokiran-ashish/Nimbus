package consultation

import (
	"context"
	"fmt"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	worflowutils "github.com/astrokiran/nimbus/internal/workflows/utils"

	"github.com/google/uuid"
)

func (con *Consultation) HandleConsultationCreation(req CreateConsultationRequest) (*model.Consultation, error) {
	consultation, err := con.CreateConsultation(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create consultation: %w", err)
	}

	workflowID := fmt.Sprintf("%s:%s", constants.ConsultationWorkflowName, consultation.ConsultationID.String())
	worflowOptions := worflowutils.GetStartWorkflowOptions(constants.ConsultationWorkflowName, con.workflowTaskQueue, workflowID)

	we, err := con.engine.Client.ExecuteWorkflow(context.Background(), worflowOptions, constants.ConsultationWorkflowName, consultation)
	if err != nil {
		return nil, fmt.Errorf("failed to start workflow: %w", err)
	}

	if err = we.Get(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("failed to run workflow: %w", err)
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
