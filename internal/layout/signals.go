package layout

type Signal struct {
	// Signals are placed at TrackNodes of boundary types only (two tracks connecting)
	Approach EdgeEnd `json:"approach"` // Each signal is positioned at a node, so with an edge end we can use it to derive which direction the signal is facing in.
}

type SignalId string
