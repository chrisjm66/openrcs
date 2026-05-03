package simulation

type Command interface {
	Apply(world *WorldState) error
}

type UpdateTrackCircuitCommand struct {
	Command
	trackCircuitId    string
	trackCircuitState TrackCircuitState
}
