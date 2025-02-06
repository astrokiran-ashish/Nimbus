package consultation

import "github.com/google/uuid"

type CreateConsultationRequest struct {
	UserID           uuid.UUID `json:"user_id"`
	ConsultantID     uuid.UUID `json:"consultant_id"`
	ConsultationType string    `json:"consultation_type"`
	AgoraChannel     string    `json:"agora_channel"`
	SessionID        uuid.UUID `json:"session_id"`
}

type UpdateConsultationRequest struct {
	ConsultationID       uuid.UUID `json:"consultation_id"`
	State                string    `json:"state"`
	UserWaitTimeSecs     int64     `json:"user_wait_time_secs"`
	ConsultationTimeSecs int64     `json:"consultation_time_secs"`
}
