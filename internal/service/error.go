package service

import (
	"github.com/chrisjm66/openrcs/internal/layout"
)

// For service bindings
func (s *SimulationNotFoundError) Error() string {
	return "Error: simulation " + string(s.Simulation) + " not found."
}

type SimulationNotFoundError struct {
	Simulation layout.RailwayLayoutId
}
