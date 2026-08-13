package state

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
