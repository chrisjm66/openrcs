package state

import "github.com/chrisjm66/openrcs/internal/layout"

func (state *WorldState) SwitchDirection (switchId layout.SwitchId) SwitchDirection {
	return state.switchState[switchId].switchDirection
}

func (state *WorldState) SetSwitchDirection(switchId layout.SwitchId, direction SwitchDirection) {
	newState := state.switchState[switchId]
	newState.switchDirection = direction
	state.switchState[switchId] = newState
}

type SwitchState struct {
	switchDirection SwitchDirection	
}

type SwitchDirection int

const (
	Normal SwitchDirection = iota
	Reverse
)
