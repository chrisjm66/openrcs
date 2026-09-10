use std::collections::HashMap;

use crate::layout::{
    SimulationLayout,
    signal::{Signal, SignalId},
    switch::{Switch, SwitchId},
    track::{
        EdgeEnd, Point, TrackEdge, TrackEdgeId, TrackNode, TrackNodeId, TrackProperties,
        TrackType::{Boundary, Switch as SwitchNode},
    },
    track_circuit::{TrackCircuit, TrackCircuitId},
};

/// A compact passing-loop layout used by the UI and simulation fixtures.
///
/// The main line and loop join at two points, allowing the scenario to exercise
/// route selection, occupancy colouring, signal aspects, and switch symbols.
pub fn make_test_layout() -> SimulationLayout {
    SimulationLayout {
        track_nodes: create_track_nodes(),
        track_edges: create_track_edges(),
        track_circuits: create_track_circuits(),
        signals: create_signals(),
        switches: create_switches(),
    }
}

fn create_track_nodes() -> HashMap<TrackNodeId, TrackNode> {
    HashMap::from([
        (
            "WEST_BOUNDARY".to_owned(),
            TrackNode {
                position: Point { x: -400.0, y: 0.0 },
                track_type: Boundary,
            },
        ),
        (
            "WEST_JUNCTION".to_owned(),
            TrackNode {
                position: Point { x: -220.0, y: 0.0 },
                track_type: SwitchNode,
            },
        ),
        (
            "EAST_JUNCTION".to_owned(),
            TrackNode {
                position: Point { x: 220.0, y: 0.0 },
                track_type: SwitchNode,
            },
        ),
        (
            "EAST_BOUNDARY".to_owned(),
            TrackNode {
                position: Point { x: 400.0, y: 0.0 },
                track_type: Boundary,
            },
        ),
    ])
}

fn create_track_edges() -> HashMap<TrackEdgeId, TrackEdge> {
    HashMap::from([
        (
            "WEST_APPROACH".to_owned(),
            TrackEdge {
                from: "WEST_BOUNDARY".to_owned(),
                to: "WEST_JUNCTION".to_owned(),
                geometry: vec![Point { x: -400.0, y: 0.0 }, Point { x: -220.0, y: 0.0 }],
                properties: TrackProperties {
                    electrified: true,
                    speed_limit: 60,
                },
                allows_from_to: true,
                allows_to_from: true,
            },
        ),
        (
            "MAIN_LINE".to_owned(),
            TrackEdge {
                from: "WEST_JUNCTION".to_owned(),
                to: "EAST_JUNCTION".to_owned(),
                geometry: vec![Point { x: -220.0, y: 0.0 }, Point { x: 220.0, y: 0.0 }],
                properties: TrackProperties {
                    electrified: true,
                    speed_limit: 80,
                },
                allows_from_to: true,
                allows_to_from: true,
            },
        ),
        (
            "PASSING_LOOP".to_owned(),
            TrackEdge {
                from: "WEST_JUNCTION".to_owned(),
                to: "EAST_JUNCTION".to_owned(),
                geometry: vec![
                    Point { x: -220.0, y: 0.0 },
                    Point {
                        x: -100.0,
                        y: 110.0,
                    },
                    Point { x: 100.0, y: 110.0 },
                    Point { x: 220.0, y: 0.0 },
                ],
                properties: TrackProperties {
                    electrified: true,
                    speed_limit: 40,
                },
                allows_from_to: true,
                allows_to_from: true,
            },
        ),
        (
            "EAST_APPROACH".to_owned(),
            TrackEdge {
                from: "EAST_JUNCTION".to_owned(),
                to: "EAST_BOUNDARY".to_owned(),
                geometry: vec![Point { x: 220.0, y: 0.0 }, Point { x: 400.0, y: 0.0 }],
                properties: TrackProperties {
                    electrified: true,
                    speed_limit: 60,
                },
                allows_from_to: true,
                allows_to_from: true,
            },
        ),
    ])
}

fn create_track_circuits() -> HashMap<TrackCircuitId, TrackCircuit> {
    HashMap::from([
        (
            "TC_WEST_APPROACH".to_owned(),
            TrackCircuit {
                edges: vec!["WEST_APPROACH".to_owned()],
            },
        ),
        (
            "TC_MAIN_LINE".to_owned(),
            TrackCircuit {
                edges: vec!["MAIN_LINE".to_owned()],
            },
        ),
        (
            "TC_PASSING_LOOP".to_owned(),
            TrackCircuit {
                edges: vec!["PASSING_LOOP".to_owned()],
            },
        ),
        (
            "TC_EAST_APPROACH".to_owned(),
            TrackCircuit {
                edges: vec!["EAST_APPROACH".to_owned()],
            },
        ),
    ])
}

fn create_signals() -> HashMap<SignalId, Signal> {
    HashMap::from([
        (
            "SIG_WEST_ENTRY".to_owned(),
            Signal {
                approach: EdgeEnd {
                    node_id: "WEST_JUNCTION".to_owned(),
                    edge_id: "WEST_APPROACH".to_owned(),
                },
            },
        ),
        (
            "SIG_WEST_EXIT".to_owned(),
            Signal {
                approach: EdgeEnd {
                    node_id: "WEST_JUNCTION".to_owned(),
                    edge_id: "MAIN_LINE".to_owned(),
                },
            },
        ),
        (
            "SIG_EAST_EXIT".to_owned(),
            Signal {
                approach: EdgeEnd {
                    node_id: "EAST_JUNCTION".to_owned(),
                    edge_id: "MAIN_LINE".to_owned(),
                },
            },
        ),
        (
            "SIG_EAST_ENTRY".to_owned(),
            Signal {
                approach: EdgeEnd {
                    node_id: "EAST_JUNCTION".to_owned(),
                    edge_id: "EAST_APPROACH".to_owned(),
                },
            },
        ),
    ])
}

fn create_switches() -> HashMap<SwitchId, Switch> {
    HashMap::from([
        (
            "SW_WEST".to_owned(),
            Switch {
                common: EdgeEnd {
                    node_id: "WEST_JUNCTION".to_owned(),
                    edge_id: "WEST_APPROACH".to_owned(),
                },
                normal: EdgeEnd {
                    node_id: "WEST_JUNCTION".to_owned(),
                    edge_id: "MAIN_LINE".to_owned(),
                },
                reverse: EdgeEnd {
                    node_id: "WEST_JUNCTION".to_owned(),
                    edge_id: "PASSING_LOOP".to_owned(),
                },
            },
        ),
        (
            "SW_EAST".to_owned(),
            Switch {
                common: EdgeEnd {
                    node_id: "EAST_JUNCTION".to_owned(),
                    edge_id: "EAST_APPROACH".to_owned(),
                },
                normal: EdgeEnd {
                    node_id: "EAST_JUNCTION".to_owned(),
                    edge_id: "MAIN_LINE".to_owned(),
                },
                reverse: EdgeEnd {
                    node_id: "EAST_JUNCTION".to_owned(),
                    edge_id: "PASSING_LOOP".to_owned(),
                },
            },
        ),
    ])
}
