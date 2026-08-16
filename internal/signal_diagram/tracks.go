package signaldiagram

import "github.com/chrisjm66/openrcs/internal/layout"

type DiagramTrack struct {
	DiagramPositions []DiagramPosition       `json:"diagramPositions"`
	TrackCircuits    []layout.TrackCircuitId `json:"trackCircuits"`
}
