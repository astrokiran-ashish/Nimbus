package consultant

import (
	"encoding/json"
	"net/http"

	common_errors "github.com/astrokiran/nimbus/internal/common/errors"
	"github.com/astrokiran/nimbus/internal/common/response"
)

func (ast *Consultant) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	session, err := ast.GenerateOTP(req.AreaCode + req.PhoneNumber)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to generate OTP", nil)
		return
	}

	err = response.JSON(w, http.StatusOK, &LoginViaOTPResponse{
		OTP:       int64(*session.Otp),
		Message:   "OTP sent successfully",
		UserID:    session.UserID.String(),
		SessionID: session.SessionID.String(),
	})
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to send OTP", nil)
		return
	}
}

func (c *Consultant) ListConsultantsHandler(w http.ResponseWriter, r *http.Request) {
	consultants, err := c.ListConsultants()
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to get consultants", nil)
		return
	}

	err = response.JSON(w, http.StatusOK, consultants)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to send response", nil)
		return
	}
}

func (c *Consultant) GetConsultant(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phone_number")
	if phoneNumber == "" {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Phone number is required", nil)
		return
	}

	consultant, err := c.GetConsultantByPhoneNumber(phoneNumber)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to get consultant", nil)
		return
	}
	if consultant == nil {
		common_errors.ErrorMessage(w, r, http.StatusNotFound, "Consultant not found", nil)
		return
	}
	response.JSON(w, http.StatusOK, consultant)
}
