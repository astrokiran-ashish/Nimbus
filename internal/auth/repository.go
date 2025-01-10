package auth

import (
	"time"

	"github.com/astrokiran/nimbus/internal/common/configs"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
	"github.com/google/uuid"
)

func (auth *Auth) CreateSession(userID uuid.UUID, phonenumber string, otp int64) error {
	session := &model.UserAuth{
		UserID:          userID,
		SessionID:       uuid.New(),
		PhoneNumber:     phonenumber,
		Otp:             int32(otp),
		OtpCreatedAt:    time.Now(),
		OtpValiditySecs: int32(configs.GetInt("OTP_VALIDITY_SECS", 60)),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	stmt := table.UserAuth.INSERT(
		table.UserAuth.UserID,
		table.UserAuth.SessionID,
		table.UserAuth.PhoneNumber,
		table.UserAuth.Otp,
		table.UserAuth.OtpCreatedAt,
		table.UserAuth.OtpValiditySecs,
		table.UserAuth.CreatedAt,
		table.UserAuth.UpdatedAt,
	).VALUES(
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
