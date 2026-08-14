package state

import "github.com/chrisjm66/openrcs/internal/layout"

func (state *WorldState) SignalAspect(signalId layout.SignalId) SignalAspect {
	return state.signalState[signalId].aspect
}

func (state *WorldState) setSignalAspect(signalId layout.SignalId, newAspect SignalAspect) {
	newState := state.signalState[signalId]
	newState.aspect = newAspect
	state.signalState[signalId] = newState
}

type SignalState struct {
	aspect SignalAspect
}

type SignalAspect int

const (
	Red SignalAspect = iota
	Yellow
	DoubleYellow
	Green
)
