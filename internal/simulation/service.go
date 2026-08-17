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
	scenarios, err := GetScenarios()

	if err != nil {
		return &service.ApiError{
			Code:    service.LAYOUT_NOT_FOUND,
			Message: "Could not load scenarios",
		}
	}

	foundIndex := -1
	for i, scenario := range scenarios {
		if scenarioId == scenario.Id {
			foundIndex = i
		}
	}

	if foundIndex == -1 {
		return &service.ApiError{
			Code:    service.LAYOUT_NOT_FOUND,
			Message: "Could not find scenario " + string(scenarioId),
		}
	}

	s.runtime = createSimRuntime(&scenarios[foundIndex], s.ctx)

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

func (s *SimulationService) GetScenarios() ([]Scenario, service.ApiError) {
	scenarios, err := GetScenarios()

	if err != nil {
		return []Scenario{}, service.ApiError{
			Code:    service.LAYOUT_NOT_FOUND,
			Message: "Could not find scenarios directory",
		}
	}

	return scenarios, service.ApiError{}
}

func (s *SimulationService) GetState() state.WorldState {
	return *s.runtime.engine.state
}
