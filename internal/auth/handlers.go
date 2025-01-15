package auth

import (
	"encoding/json"
	"net/http"

	commonErrors "github.com/astrokiran/nimbus/internal/common/errors"
	"github.com/astrokiran/nimbus/internal/common/response"
	"go.uber.org/zap"
)

func (auth *Auth) LoginViaOTP(w http.ResponseWriter, r *http.Request) {

	var req LoginViaOTPRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		commonErrors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	otp, err := auth.GenerateOTPForPhonenumber(req.AreaCode + req.PhoneNumber)
	if err != nil {
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to generate OTP", nil)
		return
	}

	err = response.JSON(w, http.StatusOK, &LoginViaOTPResponse{
		OTP:     otp,
		Message: "OTP sent successfully",
	})
	if err != nil {
		auth.logger.Error("Error while sending OTP to customer", zap.Any("err", err))
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to send OTP", nil)
		return
	}
}
