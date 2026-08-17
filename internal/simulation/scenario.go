package simulation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/chrisjm66/openrcs/internal/layout"
	signaldiagram "github.com/chrisjm66/openrcs/internal/signal_diagram"
)

type Scenario struct {
	Id          ScenarioId `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`

	Layout            layout.RailwayLayout        `json:"railwayLayout"`
	SignallingDiagram signaldiagram.SignalDiagram `json:"signallingDiagram"`
	// interlocking
}

type ScenarioId string

func GetScenarios() ([]Scenario, error) {
	entries, err := os.ReadDir("./scenarios")
	scenarios := []Scenario{}

	if err != nil {
		return []Scenario{}, errors.New("Could not view scenarios directory")
	}

	for _, entry := range entries {
		file, err := os.Open("./scenarios/" + entry.Name())

		if err != nil {
			slog.Error("Could not read file: Unable to open " + entry.Name())
			continue
		}

		bytes, err := io.ReadAll(file)

		if err != nil {
			slog.Error("Could not read file: IO input failed")
			continue
		}

		scenario := &Scenario{}
		err = json.Unmarshal(bytes, scenario)
		scenarios = append(scenarios, *scenario)
		fmt.Print(scenarios)
	}

	return scenarios, nil
}
