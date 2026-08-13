package simulation

import (
	"fmt"

	"github.com/chrisjm66/openrcs/internal/command"
)

func (engine *SimulationEngine) step(commands []command.CommandRequest) {
	for _, command := range commands {
		fmt.Printf("command: %v\n", command)
	}
}

type SimulationEngine struct {
	state WorldState
}
