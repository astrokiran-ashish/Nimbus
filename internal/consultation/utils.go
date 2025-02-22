package consultation

import (
	"fmt"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	constants "github.com/astrokiran/nimbus/internal/workflows/constants"
	workflowstates "github.com/astrokiran/nimbus/internal/workflows/workflow-states"
)

func toConsultationState(con *model.Consultation) workflowstates.ConsultationState {
	state := workflowstates.ConsultationState{
		ConsultationID:           con.ConsultationID.String(),
		UserID:                   con.UserID.String(),
		ConsultantID:             con.ConsultantID.String(),
		SessionID:                con.SessionID.String(),
		ConsultationState:        con.ConsultationState,
		ConsultationType:         con.ConsultationType,
		ConsultantPricePerMinute: 0,
		UserWalletBalance:        0,
	}
	return state
}

func EventRouter(event ConsultantActionEventRequest) (string, string, error) {
	var workflowID, signalChannelName string
	fmt.Printf("event: %v\n", event)
	switch event.Action {
	case "accept", "reject":
		workflowID = fmt.Sprintf("%s:%s", constants.ConsultationHandShakeWorkflowName, event.ConsultationID)
		signalChannelName = constants.SessionEndSignalCh
	case "on_call":
		workflowID = fmt.Sprintf("%s:%s", constants.ConsultationHandShakeWorkflowName, event.ConsultationID)
		signalChannelName = constants.ConsultantOnCallCh
	case "user_accept_call":
		workflowID = fmt.Sprintf("%s:%s", constants.ConsultationHandShakeWorkflowName, event.ConsultationID)
		signalChannelName = constants.UserResponseSignalCh
	default:
		return "", "", fmt.Errorf("invalid action")
	}
	return workflowID, signalChannelName, nil
}
