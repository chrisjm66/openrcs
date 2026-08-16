package simulation

import (
	"context"
	"github.com/chrisjm66/openrcs/internal/service"
	"github.com/chrisjm66/openrcs/internal/state"
)

type SimulationService struct {
	ctx     context.Context
	runtime *SimulationRuntime
}

func (s *SimulationService) NewRuntime(scenarioId ScenarioId) *service.ApiError {
	scenario, ok := GetTestScenarios()[scenarioId]

	if !ok {
		return &service.ApiError{
			Code:    service.LAYOUT_NOT_FOUND,
			Message: "Scenario " + string(scenarioId) + " not found",
		}
	}

	s.runtime = createSimRuntime(&scenario, s.ctx)

	if s.runtime == nil {
		return &service.ApiError{
			Code:    service.RUNTIME_NOT_FOUND,
			Message: "Simulation runtime not found",
		}
	}

	return nil
}

func (s *SimulationService) GetCurrentScenario() Scenario {
	return *s.runtime.engine.scenario
}

func (s *SimulationService) GetScenarios() map[ScenarioId]Scenario {
	return GetTestScenarios()
}

func (s *SimulationService) GetState() state.WorldState {
	return *s.runtime.engine.state
}
