package auth

import "github.com/google/uuid"

type LoginViaOTPRequest struct {
	AreaCode    string `json:"area_code"`
	PhoneNumber string `json:"phone_number"`
}
type VerifyOTPRequest struct {
	OTP       int64     `json:"otp"`
	UserID    uuid.UUID `json:"user_id"`
	SessionID uuid.UUID `json:"session_id"`
}
