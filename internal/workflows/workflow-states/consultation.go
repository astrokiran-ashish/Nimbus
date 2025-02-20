package workflowstates

type ConsultationState struct {
	ConsultationID           string `json:"consultation_id"`
	UserID                   string `json:"user_id"`
	ConsultantID             string `json:"consultant_id"`
	SessionID                string `json:"session_id"`
	ConsultationState        string `json:"consultation_state"`
	ConsultationType         string `json:"consultation_type"`
	ConsultantPricePerMinute int32  `json:"consultant_price_per_minute"`
	UserWalletBalance        int32  `json:"user_wallet_balance"`
}

type ConsultantActionEvent struct {
	ConsultationID string `json:"consultation_id"`
	Action         string `json:"action"`
}
