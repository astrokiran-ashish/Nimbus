package users

import (
	"go.uber.org/zap"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
)

func (users *Users) getUserByPhonenumber(phonenumber string) (*model.User, error) {
	stmt := table.User.SELECT(
		table.User.AllColumns,
	).FROM(table.User).WHERE(
		table.User.PhoneNumber.EQ(users.db.Dialect.String(phonenumber)),
	)

	res := model.User{}
	err := stmt.Query(users.db.Conn, &res)
	if err != nil {
		if err == qrm.ErrNoRows {
			users.logger.Info("No user found with the given phonenumber", zap.Any("phonenumber", phonenumber))
			return nil, nil
		}
		return nil, err
	}
	return &res, nil

}

func (users *Users) createUser(user *model.User) error {
	userID := uuid.New()
	user.UserID = userID
	stmt := table.User.INSERT(
		table.User.UserID,
		table.User.PhoneNumber,
		table.User.Version,
	).VALUES(
		user.UserID,
		user.PhoneNumber,
		user.Version,
	)

	_, err := stmt.Exec(users.db.Conn)
	if err != nil {
		return err
	}
	return nil
}
