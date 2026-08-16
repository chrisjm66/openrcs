package layout

type TrackNode struct {
	Position Point
	Type     NodeType
}

type TrackEdge struct {
	From, To TrackNodeId
	Geometry []Point

	Properties TrackProperties

	AllowsToFrom, AllowsFromTo bool
}

type NodeType string

const (
	NodeBoundary NodeType = "boundary"
	NodeBuffer   NodeType = "buffer"   // Buffers/track ends
	NodeSwitch   NodeType = "switch"   // Switches
	NodeCrossing NodeType = "crossing" // Intersectings tracks

)

type TrackNodeId string

type TrackEdgeId string

type TrackProperties struct {
	Name        string `json:"name"`
	SpeedLimit  float64 `json:"speedLimit"`
	Electrified bool `json:"electrified"`
}

type EdgeEnd struct {
	NodeId TrackNodeId `json:"nodeId"`
	EdgeId TrackEdgeId `json:"edgeId"`
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
