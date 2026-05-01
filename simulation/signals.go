package simulation

type Signal struct {
	id    string
	state SignalState
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
