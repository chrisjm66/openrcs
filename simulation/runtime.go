package simulation

import (
	"fmt"
	"time"
)

func CreateSimRuntime(layout *RailwayLayout) *SimulationRuntime {
	engine := SimulationEngine{
		state: *createInitalWorldState(layout),
	}

	return &SimulationRuntime{
		engine:   &engine,
		commands: make(chan CommandRequest),
		tickRate: 50 * time.Millisecond,
	}
}

func (runtime *SimulationRuntime) StartSimulation() {
	go runtime.initializeTickLoop()
}

func (runtime *SimulationRuntime) initializeTickLoop() {
	ticker := time.NewTicker(runtime.tickRate)
	defer ticker.Stop()

	for range ticker.C {
		commands := runtime.drainCommands()
		runtime.engine.step(commands)
		time.Sleep(50 * time.Millisecond)
	}
}

func (runtime *SimulationRuntime) SubmitCommand(command CommandRequest) bool {
	fmt.Println("command submitted")
	select {
	case runtime.commands <- command:
		return true
	default:
		return false
	}
}

func (runtime *SimulationRuntime) drainCommands() []CommandRequest {
	var commands []CommandRequest
	for {
		select {
		case command := <-runtime.commands:
			commands = append(commands, command)
		default:
			return commands
		}
	}
}

type SimulationRuntime struct {
	engine   *SimulationEngine
	commands chan CommandRequest
	tickRate time.Duration
}
