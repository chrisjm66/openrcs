package layout

type Signal struct {
	Protects []TrackCircuitId
	Position TrackPosition
}

type SignalId string
