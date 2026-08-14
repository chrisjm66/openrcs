package state

import "github.com/chrisjm66/openrcs/internal/layout"

func (state *WorldState) InitializeState(railwayLayout *layout.RailwayLayout) {
	signalState := make(map[layout.SignalId]SignalState)
	trackCircuitState := make(map[layout.TrackCircuitId]TrackCircuitState)
	switchState := make(map[layout.SwitchId]SwitchState)
	trainState := make(map[TrainId]TrainState)

	// Initialize signals
	for id := range railwayLayout.Signals {
		signalState[layout.SignalId(id)] = SignalState{
			aspect: Red,
		}
	}

	// Initialize track circuits
	for id := range railwayLayout.TrackCircuits {
		trackCircuitState[layout.TrackCircuitId(id)] = TrackCircuitState{
			occupied: false,
			failed:   false,
		}
	}

	// Initialize switches
	for id := range railwayLayout.Switches {
		switchState[layout.SwitchId(id)] = SwitchState{
			switchDirection: Normal,
		}
	}

	state = &WorldState{
		signalState:       signalState,
		trackCircuitState: trackCircuitState,
		switchState:       switchState,
		trainState:        trainState,
	}
}

type WorldState struct {
	signalState       map[layout.SignalId]SignalState
	trackCircuitState map[layout.TrackCircuitId]TrackCircuitState
	switchState       map[layout.SwitchId]SwitchState
	trainState        map[TrainId]TrainState
}
