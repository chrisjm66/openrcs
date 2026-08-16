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
	Name        string
	SpeedLimit  float64
	Electrified bool
}

type EdgeEnd struct {
	NodeId TrackNodeId
	EdgeId TrackEdgeId
}

type Point struct {
	X, Y float64
}
