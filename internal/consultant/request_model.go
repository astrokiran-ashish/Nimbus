package consultant

type LoginRequest struct {
	AreaCode    string `json:"area_code"`
	PhoneNumber string `json:"phone_number"`
}
