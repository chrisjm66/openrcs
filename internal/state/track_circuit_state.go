package state

import "github.com/chrisjm66/openrcs/internal/layout"

func (state *WorldState) TrackCircuitOccupied(trackCircuitId layout.TrackCircuitId) bool {
	return state.trackCircuitState[trackCircuitId].occupied
}

func (state *WorldState) OccupyTrackCircuit(trackCircuitId layout.TrackCircuitId) {
	newState := trackCircuitState[trackCircuitId]
	newState.occupied = true
	state.trackCircuitState[trackCircuitId] = newState
}

func (state *WorldState) UnoccupyTrackCircuit(trackCircuitId layout.TrackCircuitId) {
	newState := state.trackCircuitState[trackCircuitId]
	newState.occupied = false
	state.trackCircuitState[trackCircuitId] = newState
}

type TrackCircuitState struct {
	occupied bool
	failed   bool
}
