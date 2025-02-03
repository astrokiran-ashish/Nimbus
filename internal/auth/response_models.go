package auth

type LoginViaOTPResponse struct {
	OTP       int64  `json:"otp"`
	Message   string `json:"message,omitempty"`
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
}

type VerifyOTPResponse struct {
	IsValid      bool   `json:"is_valid"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Message      string `json:"message,omitempty"`
}
