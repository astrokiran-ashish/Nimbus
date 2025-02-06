package workflows

import orchestrationengine "github.com/astrokiran/nimbus/internal/common/orchestration-engine"

type Workflows struct {
	engine *orchestrationengine.OrchestrationEngine
}

func NewWorkflows(orchestrationengine *orchestrationengine.OrchestrationEngine) *Workflows {
	return &Workflows{engine: orchestrationengine}
}
