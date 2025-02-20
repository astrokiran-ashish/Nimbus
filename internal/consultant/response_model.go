package consultant

type LoginViaOTPResponse struct {
	OTP       int64  `json:"otp"`
	Message   string `json:"message,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
}

type VerifyOTPResponse struct {
	IsValid  bool   `json:"is_valid"`
	JWTToken string `json:"jwt_token,omitempty"`
	Message  string `json:"message,omitempty"`
}
