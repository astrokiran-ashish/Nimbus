package consultation

import (
	"github.com/astrokiran/nimbus/internal/common/database"
	orchestrationengine "github.com/astrokiran/nimbus/internal/common/orchestration-engine"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/google/uuid"
)

type IConsultation interface {
	CreateConsultation(req CreateConsultationRequest) (*model.Consultation, error)
	GetConsultationByID(consultationID uuid.UUID) (*model.Consultation, error)
	UpdateConsultation(req UpdateConsultationRequest) (*model.Consultation, error)
}

type Consultation struct {
	db                *database.Database
	engine            *orchestrationengine.OrchestrationEngine
	workflowTaskQueue string
}

func NewConsultation(db *database.Database, engine *orchestrationengine.OrchestrationEngine, workflowTaskQueue string) *Consultation {
	return &Consultation{
		db:                db,
		engine:            engine,
		workflowTaskQueue: workflowTaskQueue,
	}
}
