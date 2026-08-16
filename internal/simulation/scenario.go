package simulation

import (
	"github.com/chrisjm66/openrcs/internal/layout"
	signaldiagram "github.com/chrisjm66/openrcs/internal/signal_diagram"
)

type Scenario struct {
	Id          ScenarioId `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`

	Layout            layout.RailwayLayout        `json:"railwayLayout"`
	SignallingDiagram signaldiagram.SignalDiagram `json:"signallingDiagram"`
	// interlocking
}

type ScenarioId string

func GetTestScenarios() map[ScenarioId]Scenario {
	return map[ScenarioId]Scenario{
		ScenarioId("SC1"): {
			Name:        "Test Scenario 1",
			Description: "This is a test",
			Layout: layout.RailwayLayout{
				TrackNodes: map[layout.TrackNodeId]layout.TrackNode{
					layout.TrackNodeId("N1"): {
						Position: layout.Point{
							X: 100,
							Y: 100,
						},
						Type: layout.NodeBuffer,
					},
					layout.TrackNodeId("N3"): {
						Position: layout.Point{
							X: 200,
							Y: 300,
						},
						Type: layout.NodeBoundary,
					},
					layout.TrackNodeId("N2"): {
						Position: layout.Point{
							X: 400,
							Y: 100,
						},
						Type: layout.NodeBuffer,
					},
				},
				TrackEdges: map[layout.TrackEdgeId]layout.TrackEdge{
					layout.TrackEdgeId("TE1"): {
						Properties: layout.TrackProperties{
							Name:        "idk",
							Electrified: false,
							SpeedLimit:  60,
						},
						Geometry: []layout.Point{
							{
								X: 150,
								Y: 200,
							},
						},
					},
					layout.TrackEdgeId("TE2"): {
						Properties: layout.TrackProperties{
							Name:        "idk",
							Electrified: false,
							SpeedLimit:  60,
						},
						Geometry: []layout.Point{
							{
								X: 150,
								Y: 200,
							},
						},
					},
				},
				TrackCircuits: map[layout.TrackCircuitId]layout.TrackCircuit{
					layout.TrackCircuitId("TC1"): {
						Edges: []layout.TrackEdgeId{layout.TrackEdgeId("TE1")},
					},
					layout.TrackCircuitId("TC2"): {
						Edges: []layout.TrackEdgeId{layout.TrackEdgeId("TE2")},
					},
				},
				Signals: map[layout.SignalId]layout.Signal{
					layout.SignalId("S001"): {
						Approach: layout.EdgeEnd{
							NodeId: layout.TrackNodeId("N1"),
							EdgeId: layout.TrackEdgeId("E1"),
						},
					},
				},
			},
			SignallingDiagram: signaldiagram.SignalDiagram{
				Tracks: []signaldiagram.DiagramTrack{
					{
						DiagramPositions: []signaldiagram.DiagramPosition{
							{
								X: 100,
								Y: 100,
							},
							{
								X: 150,
								Y: 100,
							},
							{
								X: 1000,
								Y: 150,
							},
						},
						TrackCircuits: []layout.TrackCircuitId{
							"TC1",
						},
					},
				},
				Signals: []signaldiagram.DiagramSignal{
					{
						SignalId: "S001",
						DiagramPosition: signaldiagram.DiagramPosition{
							X: 100,
							Y: 500,
						},
						Orientation: 90,
					},
				},
			},
		},
	}
}
