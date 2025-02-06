package workflowstates

import "github.com/google/uuid"

type ConsultationState struct {
	ConsultationID    string    `json:"consultation_id"`
	UserID            string    `json:"user_id"`
	ConsultantID      uuid.UUID `json:"consultant_id"`
	SessionID         string    `json:"session_id"`
	ConsultationState string    `json:"consultation_state"`
	ConsultationType  string    `json:"consultation_type"`
}
