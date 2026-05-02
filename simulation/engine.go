package simulation

import "fmt"

func (engine *SimulationEngine) step(commands []Command) {
	for _, command := range commands {
		fmt.Printf("command: %v\n", command)
		command.Apply(&engine.state)
	}
	fmt.Println("Stepped")
}

type SimulationEngine struct {
	state WorldState
}
