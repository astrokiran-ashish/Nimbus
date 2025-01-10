package auth

import "github.com/astrokiran/nimbus/internal/models/nimbus/public/model"

func (auth *Auth) generateOTPForPhonenumber(phoneNumber string) (int64, error) {

	// Get User matching the phone number
	user, err := auth.Users.GetOrCreateUser(&model.User{
		PhoneNumber: &phoneNumber,
	})
	if err != nil {
		return 0, err
	}

	// Generate OTP
	otp := auth.generateRandomSixDigit()

	// Create Session
	err = auth.CreateSession(user.UserID, phoneNumber, otp)
	if err != nil {
		return 0, err
	}

	// Send OTP via SMS
	err = auth.SMSService.SendOTP(phoneNumber, otp)
	if err != nil {
		return 0, err
	}

	return otp, nil
}
