package simulation

func CreateDefaultTrackCircuit() TrackCircuit {
	return TrackCircuit{
		state: TrackCircuitState{
			occupied: false,
			failed:   false,
		},
	}
}

type TrackCircuit struct {
	state TrackCircuitState
}

type TrackCircuitState struct {
	occupied bool
	failed   bool
}
