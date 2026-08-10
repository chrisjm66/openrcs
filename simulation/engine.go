package simulation

import "fmt"

func (engine *SimulationEngine) step(commands []CommandRequest) {
	for _, command := range commands {
		fmt.Printf("command: %v\n", command)
	}
}

type SimulationEngine struct {
	state WorldState
}
