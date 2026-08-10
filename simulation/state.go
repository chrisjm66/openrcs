package simulation

func createInitalWorldState(layout *RailwayLayout) *WorldState {
	signalState := make(map[string]SignalState)
	trackCircuitState := make(map[string]TrackCircuitState)

	for id, _ := range layout.signals {
		signalState[id] = SignalState{
			aspect: Red,
		}
	}

	for id, _ := range layout.trackCircuits {
		trackCircuitState[id] = TrackCircuitState{
			occupied: false,
			failed:   false,
		}
	}

	worldState := WorldState{
		signalState:       signalState,
		trackCircuitState: trackCircuitState,
	}

	return &worldState
}

type WorldState struct {
	signalState       map[string]SignalState
	trackCircuitState map[string]TrackCircuitState
}
