package state

import "github.com/chrisjm66/openrcs/internal/layout"

type TrainId string

type TrainDef struct {
	TrainId      TrainId
	Description  string
	Acceleration float64
	Braking      float64
}

type TrainPosition struct {
	ReferenceEdge layout.EdgeEnd
	Offset        float64
}

type TrainState struct {
	headcode string
	speed    float64
	position TrainPosition
	length   float64
}
