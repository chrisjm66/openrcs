package layout

func CreateDefaultTrackCircuit() TrackCircuit {
	return TrackCircuit{}
}

type TrackCircuit struct {
	StartPosition TrackPosition
	EndPosition   TrackPosition
}

type TrackCircuitId string
