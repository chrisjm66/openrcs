package signaldiagram

import "github.com/chrisjm66/openrcs/internal/layout"

type SignalDiagram struct {
	Tracks   []DiagramTrack                    `json:"tracks"`
	Signals  []DiagramSignal                   `json:"signals"`
	Switches map[layout.SwitchId]DiagramSwitch `json:"switches"`
}

type DiagramPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}
