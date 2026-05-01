package simulation

type WorldState struct {
	signalState       map[string]SignalState
	trackCircuitState map[string]TrackCircuitState
}
