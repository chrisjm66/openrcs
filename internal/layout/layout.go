package layout

type RailwayLayoutId string

type RailwayLayout struct {
	TrackNodes    map[TrackNodeId]TrackNode
	TrackEdges    map[TrackEdgeId]TrackEdge
	Signals       map[SignalId]Signal
	TrackCircuits map[TrackCircuitId]TrackCircuit
	Switches      map[SwitchId]Switch
}
