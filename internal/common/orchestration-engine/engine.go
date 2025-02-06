package orchestrationengine

import "go.temporal.io/sdk/client"

type OrchestrationEngine struct {
	Client client.Client
}

func NewOrchestrationEngine(c client.Client) *OrchestrationEngine {
	return &OrchestrationEngine{Client: c}
}
