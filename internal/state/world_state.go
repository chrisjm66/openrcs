package state

import "github.com/chrisjm66/openrcs/internal/layout"

func CreateInitalWorldState(railwayLayout *layout.RailwayLayout) *WorldState {
	signalState := make(map[layout.SignalId]SignalState)
	trackCircuitState := make(map[layout.TrackCircuitId]TrackCircuitState)

	for id := range railwayLayout.Signals {
		signalState[layout.SignalId(id)] = SignalState{
			aspect: Red,
		}
	}

	for id := range railwayLayout.TrackCircuits {
		trackCircuitState[layout.TrackCircuitId(id)] = TrackCircuitState{
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
	signalState       map[layout.SignalId]SignalState
	trackCircuitState map[layout.TrackCircuitId]TrackCircuitState
	switchState map[layout.SwitchId]SwitchState
}
