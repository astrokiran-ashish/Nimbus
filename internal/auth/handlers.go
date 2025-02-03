package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func (auth *Auth) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req VerifyOTPRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		commonErrors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	session, err := auth.GetSession(req.UserID, req.SessionID)
	if err != nil {
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to retrieve session", nil)
		return
	}
	fmt.Println("time.now", time.Now().Local())
	fmt.Println("otp time", session.OtpCreatedAt.Add(time.Duration(*session.OtpValiditySecs)*time.Second).Local())
	fmt.Println("Time After", time.Now().After(session.OtpCreatedAt.Add(time.Duration(*session.OtpValiditySecs)*time.Second)))
	if time.Now().After(session.OtpCreatedAt.Add(time.Duration(*session.OtpValiditySecs) * time.Second)) {
		commonErrors.ErrorMessage(w, r, http.StatusUnauthorized, "OTP expired", nil)
		return
	}

	if int32(*session.Otp) != int32(req.OTP) {
		commonErrors.ErrorMessage(w, r, http.StatusUnauthorized, "Invalid OTP", nil)
		return
	}

	// Generate tokens after OTP verification
	accessToken, refreshToken, err := auth.GenerateTokens(session.UserID)
	if err != nil {
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to generate tokens", nil)
		return
	}

	err = response.JSON(w, http.StatusOK, VerifyOTPResponse{
		IsValid:      true,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		auth.logger.Error("Error while sending verification response", zap.Any("err", err))
		commonErrors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to send response", nil)
		return
	}
}
