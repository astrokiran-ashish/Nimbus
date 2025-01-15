package consultant

import (
	"errors"
	"fmt"

	common_utils "github.com/astrokiran/nimbus/internal/common/utils"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
)

func (c *Consultant) GenerateOTP(phoneNumber string) (int64, error) {

	// Check if the user exists in the database
	user, err := c.user.GetOrCreateUser(&model.User{
		PhoneNumber: &phoneNumber,
	})
	if err != nil {
		return 0, err
	}
	// User not found, create a new use
	if user == nil {
		return 0, errors.New("User not found")
	}
	err = c.GetOrCreateConsultant(user.UserID)
	if err != nil {
		return 0, err
	}

	otp := common_utils.GenerateRandomSixDigit()

	// Send OTP to the user's phone number
	// Create Session
	err = c.auth.CreateSession(user.UserID, phoneNumber, otp)
	if err != nil {
		return 0, err
	}

	// Send OTP via SMS
	err = c.auth.SMSService.SendOTP(phoneNumber, otp)
	if err != nil {
		fmt.Println("Error sending OTP:", err)
		return 0, err
	}

	return otp, nil

}
