package simulation

import (
	"context"

	"github.com/chrisjm66/openrcs/internal/layout"
)

type SimulationService struct {
	ctx     context.Context
	runtime *SimulationRuntime
}

func (s *SimulationService) NewRuntime(layout layout.RailwayLayout) {
	s.runtime = createSimRuntime(&layout, s.ctx)

}

func (s *SimulationService) GetLayouts() []layout.RailwayLayoutId {
	return layout.GetAvailableLayouts()
}
