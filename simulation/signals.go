package simulation

func CreateSignal(protectsTrackCircuits []string) Signal {
	return Signal{
		protects: protectsTrackCircuits,
		state: SignalState{
			aspect: SignalAspect(Red),
		},
	}
}

type Signal struct {
	protects []string
	state    SignalState
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
