package communication

import (
	"github.com/astrokiran/nimbus/internal/common/database"
	orchestrationengine "github.com/astrokiran/nimbus/internal/common/orchestration-engine"
	"go.uber.org/zap"
)

type Communication struct {
	db     *database.Database
	logger *zap.Logger
	engine *orchestrationengine.OrchestrationEngine
}

func NewCommunication(db *database.Database, logger *zap.Logger, engine *orchestrationengine.OrchestrationEngine) *Communication {
	return &Communication{
		db:     db,
		logger: logger,
		engine: engine,
	}
}
