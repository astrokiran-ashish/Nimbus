package auth

type LoginViaOTPRequest struct {
	AreaCode    string `json:"area_code"`
	PhoneNumber string `json:"phone_number"`
}
type VerifyOTPRequest struct {
	AreaCode    string `json:"area_code"`
	PhoneNumber string `json:"phone_number"`
	OTP         int64  `json:"otp"`
}
