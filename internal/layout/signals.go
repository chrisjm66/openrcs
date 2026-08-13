package layout

func CreateSignal(protectsTrackCircuits []string) Signal {
	return Signal{
		protects: protectsTrackCircuits,
	}
}

type Signal struct {
	protects []string
}

type SignalId string
