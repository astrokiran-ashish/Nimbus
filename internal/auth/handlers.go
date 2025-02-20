package auth

import (
	"encoding/json"
	"fmt"
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

	session, err := auth.GenerateOTPForPhonenumber(req.AreaCode + req.PhoneNumber)
	if err != nil {
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to generate OTP", nil)
		return
	}

	err = response.JSON(w, http.StatusOK, &LoginViaOTPResponse{
		OTP:       int64(*session.Otp),
		Message:   "OTP sent successfully",
		UserID:    session.UserID.String(),
		SessionID: session.SessionID.String(),
	})
	if err != nil {
		auth.logger.Error("Error while sending OTP to customer", zap.Any("err", err))
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to send OTP", nil)
		return
	}
}

// VerifyOTP decodes the request, calls the service method for OTP verification, and writes the response.
func (auth *Auth) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req VerifyOTPRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		commonErrors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	verifyResp, err := auth.VerifyOTPService(req)
	if err != nil {
		// Use unauthorized error for OTP expiration or mismatch cases.
		if err.Error() == "OTP expired" || err.Error() == "invalid OTP" {
			commonErrors.ErrorMessage(w, r, http.StatusUnauthorized, err.Error(), nil)
		} else {
			commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to verify OTP", nil)
		}
		return
	}

	err = response.JSON(w, http.StatusOK, verifyResp)
	if err != nil {
		auth.logger.Error("Error while sending verification response", zap.Any("err", err))
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to send response", nil)
		return
	}
}

func (auth *Auth) VerifyToken(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ProcessToken(r)
	if err != nil {
		auth.logger.Error("Token validation failed", zap.Any("err", err))
		commonErrors.ErrorMessage(w, r, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	err = response.JSON(w, http.StatusOK, map[string]interface{}{
		"userID": userID,
	})
	if err != nil {
		auth.logger.Error("Failed to send response", zap.Any("err", err))
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to send response", nil)
		return
	}
}
