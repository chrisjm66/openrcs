package simulation

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/chrisjm66/openrcs/internal/command"
	"github.com/chrisjm66/openrcs/internal/layout"
	"github.com/chrisjm66/openrcs/internal/state"
)

func createSimRuntime(layout *layout.RailwayLayout, ctx context.Context) *SimulationRuntime {
	engine := SimulationEngine{
		state:  state.WorldState{},
		layout: *layout,
	}

	engine.state.InitializeState(layout)

	return &SimulationRuntime{
		engine:   &engine,
		commands: make(chan command.CommandRequest),
		tickRate: 50 * time.Millisecond,
		ctx:      ctx,
	}
}

func (runtime *SimulationRuntime) StartSimulation() {
	go runtime.initializeTickLoop()
}

func (runtime *SimulationRuntime) initializeTickLoop() {
	ticker := time.NewTicker(runtime.tickRate)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-runtime.ctx.Done():
			slog.Debug("Stopping runtime - stop message received from program")
			break
		default:
			commands := runtime.drainCommands()
			runtime.engine.step(commands)
			time.Sleep(50 * time.Millisecond)
		}
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
	ctx      context.Context
}
