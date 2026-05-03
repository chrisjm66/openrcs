package main

import (
	"embed"
	"openrcs/simulation"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

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

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "openrcs",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			runtime,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
