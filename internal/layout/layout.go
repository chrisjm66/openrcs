package layout

import "maps"

func CreateRailwayLayout(signals map[string]Signal, trackCircuits map[string]TrackCircuit) RailwayLayout {
	layout := RailwayLayout{
		maps.Clone(signals),
		maps.Clone(trackCircuits),
	}

	return layout
}

type RailwayLayout struct {
	Signals       map[string]Signal
	TrackCircuits map[string]TrackCircuit
}
