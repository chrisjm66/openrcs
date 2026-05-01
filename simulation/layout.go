package simulation

func CreateRailwayLayout(signals map[string]Signal, trackCircuits map[string]TrackCircuit) RailwayLayout {
	layout := RailwayLayout{signals, trackCircuits}

	return layout
}

type RailwayLayout struct {
	signals       map[string]Signal
	trackCircuits map[string]TrackCircuit
}
