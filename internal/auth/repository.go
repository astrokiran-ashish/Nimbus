package auth

import (
	"time"

	"github.com/astrokiran/nimbus/internal/common/configs"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
	"github.com/google/uuid"
)

func (auth *Auth) CreateSession(userID uuid.UUID, phonenumber string, otp int64) (*model.UserAuth, error) {
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
		return nil, err
	}

	return session, nil

}

func (a *Auth) GetSession(userId uuid.UUID, sessionId uuid.UUID) (*model.UserAuth, error) {

	stmt := table.UserAuth.SELECT(table.UserAuth.AllColumns).FROM(table.UserAuth).WHERE(
		table.UserAuth.UserID.EQ(a.db.Dialect.UUID(userId)).AND(table.UserAuth.SessionID.EQ(a.db.Dialect.UUID(sessionId))))
	session := model.UserAuth{}
	err := stmt.Query(a.db.Conn, &session)
	return &session, err
}

func (auth *Auth) UpdateSession(session *model.UserAuth) (*model.UserAuth, error) {
	now := time.Now()

	// Prepare the update statement
	stmt := table.UserAuth.UPDATE(
		table.UserAuth.JwtTokenHash,
		table.UserAuth.RefreshTokenHash,
		table.UserAuth.UpdatedAt,
	).SET(
		session.JwtTokenHash,
		session.RefreshTokenHash,
		now,
	).WHERE(
		table.UserAuth.UserID.EQ(auth.db.Dialect.UUID(session.UserID)).AND(table.UserAuth.SessionID.EQ(auth.db.Dialect.UUID(session.SessionID))),
	)

	// Execute the update statement
	_, err := stmt.Exec(auth.db.Conn)
	if err != nil {
		return nil, err
	}

	// Retrieve the updated session
	return session, nil
}
