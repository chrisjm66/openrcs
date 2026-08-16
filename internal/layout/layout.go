package layout

type RailwayLayoutId string

type RailwayLayout struct {
	TrackNodes    map[TrackNodeId]TrackNode `json:"trackNodes"`
	TrackEdges    map[TrackEdgeId]TrackEdge `json:"trackEdges"`
	Signals       map[SignalId]Signal `json:"signals"`
	TrackCircuits map[TrackCircuitId]TrackCircuit `json:"trackCircuits"`
	Switches      map[SwitchId]Switch `json:"switches"`
}
