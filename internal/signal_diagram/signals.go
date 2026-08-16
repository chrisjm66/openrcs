package signaldiagram

import "github.com/chrisjm66/openrcs/internal/layout"

type DiagramSignal struct {
	SignalId        layout.SignalId `json:"signalId"`
	DiagramPosition `json:"diagramPosition"`
	Orientation     int `json:"orientation"`
}
