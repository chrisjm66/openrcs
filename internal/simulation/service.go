package simulation

import (
	"context"
	"errors"

	"github.com/chrisjm66/openrcs/internal/layout"
	"github.com/chrisjm66/openrcs/internal/service"
	"github.com/chrisjm66/openrcs/internal/state"
)

type SimulationService struct {
	ctx     context.Context
	runtime *SimulationRuntime
}

func (s *SimulationService) NewRuntime(railwayLayoutId layout.RailwayLayoutId) error {
	simLayout, err := layout.GetLayout(railwayLayoutId)

	if err != nil {
		return &service.SimulationNotFoundError{
			Simulation: railwayLayoutId,
		}
	}

	s.runtime = createSimRuntime(&simLayout, s.ctx)

	if s.runtime == nil {
		return errors.New("Runtime not loaded")
	}

	return nil
}

func (s *SimulationService) GetLayouts() []layout.RailwayLayoutId {
	return layout.GetAvailableLayouts()
}

func (s *SimulationService) GetLayout() layout.RailwayLayout {
	return s.runtime.engine.layout
	}

func (s *SimulationService) GetState() state.WorldState {
	return s.runtime.engine.state
}
