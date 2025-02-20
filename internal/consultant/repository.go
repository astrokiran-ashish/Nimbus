package consultant

import (
	"fmt"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
)

func (c *Consultant) GetConsultantByUserID(userID uuid.UUID) (*model.Consultant, error) {
	stmt := table.Consultant.SELECT(
		table.Consultant.AllColumns,
	).FROM(table.Consultant).WHERE(
		table.Consultant.UserID.EQ(c.db.Dialect.UUID(userID)),
	)

	res := model.Consultant{}
	err := stmt.Query(c.db.Conn, &res)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
func (c *Consultant) CreateConsultant(consultant *model.Consultant) error {
	consultantID := uuid.New()
	consultant.ConsultantID = consultantID
	var version int32 = 1
	consultant.Version = &version
	stmt := table.Consultant.INSERT(
		table.Consultant.ConsultantID,
		table.Consultant.UserID,
		table.Consultant.Version,
		table.Consultant.State,
		table.Consultant.CallChannel,
		table.Consultant.ChatChannel,
		table.Consultant.VideoCallChannel,
		table.Consultant.LiveChannel,
	).VALUES(
		consultant.ConsultantID,
		consultant.UserID,
		consultant.Version,
		"Pending",
		fmt.Sprintf("%s_%s", consultant.ConsultantID.String(), "Call"),
		fmt.Sprintf("%s_%s", consultant.ConsultantID.String(), "Chat"),
		fmt.Sprintf("%s_%s", consultant.ConsultantID.String(), "VideoCall"),
		fmt.Sprintf("%s_%s", consultant.ConsultantID.String(), "Live"),
	)

	_, err := stmt.Exec(c.db.Conn)
	if err != nil {
		return err
	}
	return nil
}

func (c *Consultant) GetOrCreateConsultant(userID uuid.UUID) error {
	consultant, err := c.GetConsultantByUserID(userID)
	if err != nil {
		return err
	}
	if consultant != nil {
		return nil
	}

	err = c.CreateConsultant(&model.Consultant{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Consultant) ListConsultants() ([]*model.Consultant, error) {
	stmt := table.Consultant.SELECT(
		table.Consultant.AllColumns,
	).FROM(
		table.Consultant.INNER_JOIN(
			table.User,
			table.User.UserID.EQ(table.Consultant.UserID),
		),
	)

	var consultants []*model.Consultant
	err := stmt.Query(c.db.Conn, &consultants)
	if err != nil {
		return nil, err
	}
	return consultants, nil
}

func (c *Consultant) GetConsultantByPhoneNumber(phoneNumber string) (*model.Consultant, error) {
	stmt := table.Consultant.SELECT(
		table.Consultant.AllColumns,
	).FROM(
		table.Consultant.INNER_JOIN(
			table.User,
			table.User.UserID.EQ(table.Consultant.UserID),
		),
	).WHERE(
		table.User.PhoneNumber.EQ(c.db.Dialect.String(phoneNumber)),
	)

	res := model.Consultant{}
	fmt.Printf("stmt: %v\n", stmt.DebugSql())
	err := stmt.Query(c.db.Conn, &res)
	if err != nil {
		if err == qrm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
