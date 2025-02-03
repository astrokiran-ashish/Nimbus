package auth

import (
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"go.uber.org/zap"
)

func (auth *Auth) GenerateOTPForPhonenumber(phoneNumber string) (*model.UserAuth, error) {

	// Get User matching the phone number
	user, err := auth.Users.GetOrCreateUser(&model.User{
		PhoneNumber: &phoneNumber,
	})
	if err != nil {
		auth.logger.Error("Error while getting user", zap.Any("err", err))
		return nil, err
	}

	// Generate OTP
	otp := auth.generateRandomSixDigit()

	// Create Session
	session, err := auth.CreateSession(user.UserID, phoneNumber, otp)
	if err != nil {
		return nil, err
	}

	// Send OTP via SMS
	// err = auth.SMSService.SendOTP(phoneNumber, otp)
	// if err != nil {
	// 	fmt.Println("Error sending OTP:", err)
	// 	return nil, err
	// }

	return session, nil
}
