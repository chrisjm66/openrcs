package simulation

import "maps"

func CreateRailwayLayout(signals map[string]Signal, trackCircuits map[string]TrackCircuit) RailwayLayout {
	layout := RailwayLayout{
		maps.Clone(signals),
		maps.Clone(trackCircuits),
	}

	return layout
}

type RailwayLayout struct {
	signals       map[string]Signal
	trackCircuits map[string]TrackCircuit
}
