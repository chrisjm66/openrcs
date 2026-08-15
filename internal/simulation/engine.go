package simulation

import (
	"fmt"

	"github.com/chrisjm66/openrcs/internal/command"
	"github.com/chrisjm66/openrcs/internal/layout"
	"github.com/chrisjm66/openrcs/internal/state"
)

func (engine *SimulationEngine) step(commands []command.CommandRequest) {
	for _, command := range commands {
		fmt.Printf("command: %v\n", command)
	}
}

type SimulationEngine struct {
	state state.WorldState
	layout layout.RailwayLayout
}
