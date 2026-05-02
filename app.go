package main

import (
	"context"
	"fmt"
	"openrcs/simulation"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	layout := simulation.CreateRailwayLayout(
		map[string]simulation.Signal{
			"1": simulation.CreateSignal([]string{"1"}),
			"2": simulation.CreateSignal([]string{"2"}),
			"3": simulation.CreateSignal([]string{"3"}),
		},
		map[string]simulation.TrackCircuit{
			"1": simulation.CreateDefaultTrackCircuit(),
			"2": simulation.CreateDefaultTrackCircuit(),
			"3": simulation.CreateDefaultTrackCircuit(),
		},
	)

	runtime := simulation.CreateSimRuntime(&layout)
	runtime.StartSimulation()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
