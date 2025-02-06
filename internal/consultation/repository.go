package consultation

import (
	"fmt"
	"time"

	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/table"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

func (con *Consultation) CreateConsultation(req CreateConsultationRequest) (*model.Consultation, error) {
	currentTime := time.Now()
	consultation := &model.Consultation{
		ConsultationID:    uuid.New(),
		UserID:            req.UserID,
		ConsultantID:      req.ConsultantID,
		ConsultationType:  req.ConsultationType,
		AgoraChannel:      &req.AgoraChannel,
		CreatedAt:         &currentTime,
		UpdatedAt:         &currentTime,
		ConsultationState: ConsultationCreated,
	}
	stmt := table.Consultation.INSERT(
		table.Consultation.ConsultationID,
		table.Consultation.UserID,
		table.Consultation.ConsultantID,
		table.Consultation.ConsultationType,
		table.Consultation.AgoraChannel,
		table.Consultation.CreatedAt,
		table.Consultation.UpdatedAt,
		table.Consultation.ConsultationState,
	).VALUES(
		consultation.ConsultationID,
		consultation.UserID,
		consultation.ConsultantID,
		consultation.ConsultationType,
		consultation.AgoraChannel,
		consultation.CreatedAt,
		consultation.UpdatedAt,
		consultation.ConsultationState,
	)

	// Execute the insert statement
	_, err := stmt.Exec(con.db.Conn)
	if err != nil {
		return nil, err
	}

	return consultation, nil
}

func (con *Consultation) GetConsultation(consultationID uuid.UUID) (*model.Consultation, error) {
	stmt := table.Consultation.SELECT(
		table.Consultation.ConsultationID,
		table.Consultation.UserID,
		table.Consultation.ConsultantID,
		table.Consultation.ConsultationType,
		table.Consultation.AgoraChannel,
		table.Consultation.CreatedAt,
		table.Consultation.UpdatedAt,
		table.Consultation.ConsultationState,
	).WHERE(
		table.Consultation.ConsultationID.EQ(con.db.Dialect.UUID(consultationID)),
	)

	consultation := model.Consultation{}
	err := stmt.Query(con.db.Conn, &consultation)
	if err != nil {
		return nil, err
	}

	return &consultation, nil
}

func (con *Consultation) UpdateConsultation(req UpdateConsultationRequest) (*model.Consultation, error) {
	// Ensure ConsultationID is provided
	if req.ConsultationID == uuid.Nil {
		return nil, fmt.Errorf("ConsultationID is required")
	}

	// Initialize a consultation model instance for updates
	consultation := model.Consultation{}

	// Create an empty list of columns to update
	columnList := postgres.ColumnList{}

	// Append only non-nil fields to the update statement
	if req.State != "" {
		consultation.ConsultationState = req.State
		columnList = append(columnList, table.Consultation.ConsultationState)
	}
	waitTime := int32(req.UserWaitTimeSecs)
	if req.UserWaitTimeSecs != 0 {
		consultation.UserWaitTimeSecs = &waitTime
		columnList = append(columnList, table.Consultation.UserWaitTimeSecs)
	}
	consultationTime := int32(req.ConsultationTimeSecs)
	if req.ConsultationTimeSecs != 0 {
		consultation.ConsultationTimeSecs = &consultationTime
		columnList = append(columnList, table.Consultation.ConsultationTimeSecs)
	}

	// If there are no updates, return early
	if len(columnList) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}
	currentTime := time.Now()
	consultation.UpdatedAt = &currentTime
	columnList = append(columnList, table.Consultation.UpdatedAt)

	// Generate the update query using ColumnList and MODEL
	updateStmt := table.ConsultationTable.UPDATE(*table.Consultation).
		SET(columnList).
		MODEL(consultation).
		WHERE(table.Consultation.ConsultationID.EQ(con.db.Dialect.UUID(req.ConsultationID)))

	// Execute the update query
	_, err := updateStmt.Exec(con.db.Conn)
	if err != nil {
		return nil, fmt.Errorf("failed to update consultation: %w", err)
	}

	return &consultation, nil

}
