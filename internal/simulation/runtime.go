package simulation

import (
	"fmt"
	"time"

	"github.com/chrisjm66/openrcs/internal/command"
	"github.com/chrisjm66/openrcs/internal/layout"
	"github.com/chrisjm66/openrcs/internal/state"
)

func CreateSimRuntime(layout *layout.RailwayLayout) *SimulationRuntime {
	engine := SimulationEngine{
		state: *state.CreateInitalWorldState(layout),
	}

	return &SimulationRuntime{
		engine:   &engine,
		commands: make(chan command.CommandRequest),
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

func (runtime *SimulationRuntime) SubmitCommand(command command.CommandRequest) bool {
	fmt.Println("command submitted")
	select {
	case runtime.commands <- command:
		return true
	default:
		return false
	}
}

func (runtime *SimulationRuntime) drainCommands() []command.CommandRequest {
	var commands []command.CommandRequest
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
	commands chan command.CommandRequest
	tickRate time.Duration
}
