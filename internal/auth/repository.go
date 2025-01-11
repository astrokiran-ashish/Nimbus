package auth

import (
	"time"

	"github.com/astrokiran/nimbus/internal/common/configs"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
	"github.com/google/uuid"
)

func (auth *Auth) CreateSession(userID uuid.UUID, phonenumber string, otp int64) error {
	otpInt32 := int32(otp)
	now := time.Now()
	validitySecs := int32(configs.GetInt("OTP_VALIDITY_SECS", 60))

	session := &model.UserAuth{
		ID:              uuid.New(),
		UserID:          userID,
		SessionID:       uuid.New(),
		PhoneNumber:     phonenumber,
		Otp:             &otpInt32,
		OtpCreatedAt:    &now,
		OtpValiditySecs: &validitySecs,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	stmt := table.UserAuth.INSERT(
		table.UserAuth.ID,
		table.UserAuth.UserID,
		table.UserAuth.SessionID,
		table.UserAuth.PhoneNumber,
		table.UserAuth.Otp,
		table.UserAuth.OtpCreatedAt,
		table.UserAuth.OtpValiditySecs,
		table.UserAuth.CreatedAt,
		table.UserAuth.UpdatedAt,
	).VALUES(
		session.ID,
		session.UserID,
		session.SessionID,
		session.PhoneNumber,
		session.Otp,
		session.OtpCreatedAt,
		session.OtpValiditySecs,
		session.CreatedAt,
		session.UpdatedAt,
	)

	_, err := stmt.Exec(auth.db.Conn)
	if err != nil {
		return err
	}

	return nil

}
