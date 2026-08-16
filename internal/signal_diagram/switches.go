package signaldiagram

import "github.com/chrisjm66/openrcs/internal/layout"

type DiagramSwitch struct {
	SwitchId        layout.SwitchId `json:"switchId"`
	DiagramPosition `json:"diagramPosition"`
}
