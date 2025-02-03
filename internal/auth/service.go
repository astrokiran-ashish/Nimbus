package auth

import (
	"errors"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

// VerifyToken verifies the JWT token and returns the user ID if valid.
func (auth *Auth) ValidateToken(tokenString string) (uuid.UUID, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(auth.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		auth.logger.Error("Error validating token", zap.Any("err", err))
		return uuid.Nil, errors.New("invalid token")
	}

	return claims.UserID, nil
}
