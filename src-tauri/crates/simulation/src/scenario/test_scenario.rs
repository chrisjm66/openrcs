use std::collections::HashMap;

use crate::{
    diagram::{DiagramPosition, DiagramSignal, DiagramSwitch, DiagramTrack, SignalDiagram},
    layout::test_layout::make_test_layout,
    scenario::Scenario,
};

/// A two-point passing loop. It is deliberately small enough to understand at
/// a glance, while exercising every currently modelled diagram object.
pub fn make_test_scenario() -> Scenario {
    Scenario {
        id: "passing-loop".to_owned(),
        name: "Riverton Passing Loop".to_owned(),
        description: "A two-way single-track section with a main line, passing loop, four signals, and two points."
            .to_owned(),
        layout: make_test_layout(),
        diagram: SignalDiagram {
            tracks: vec![
                DiagramTrack {
                    positions: vec![
                        DiagramPosition { x: -400.0, y: 0.0 },
                        DiagramPosition { x: -220.0, y: 0.0 },
                    ],
                    track_circuits: vec!["TC_WEST_APPROACH".to_owned()],
                },
                DiagramTrack {
                    positions: vec![
                        DiagramPosition { x: -220.0, y: 0.0 },
                        DiagramPosition { x: 220.0, y: 0.0 },
                    ],
                    track_circuits: vec!["TC_MAIN_LINE".to_owned()],
                },
                DiagramTrack {
                    positions: vec![
                        DiagramPosition { x: -220.0, y: 0.0 },
                        DiagramPosition { x: -100.0, y: 110.0 },
                        DiagramPosition { x: 100.0, y: 110.0 },
                        DiagramPosition { x: 220.0, y: 0.0 },
                    ],
                    track_circuits: vec!["TC_PASSING_LOOP".to_owned()],
                },
                DiagramTrack {
                    positions: vec![
                        DiagramPosition { x: 220.0, y: 0.0 },
                        DiagramPosition { x: 400.0, y: 0.0 },
                    ],
                    track_circuits: vec!["TC_EAST_APPROACH".to_owned()],
                },
            ],
            signals: vec![
                DiagramSignal {
                    signal_id: "SIG_WEST_ENTRY".to_owned(),
                    position: DiagramPosition { x: -245.0, y: -18.0 },
                },
                DiagramSignal {
                    signal_id: "SIG_WEST_EXIT".to_owned(),
                    position: DiagramPosition { x: -195.0, y: 18.0 },
                },
                DiagramSignal {
                    signal_id: "SIG_EAST_EXIT".to_owned(),
                    position: DiagramPosition { x: 195.0, y: -18.0 },
                },
                DiagramSignal {
                    signal_id: "SIG_EAST_ENTRY".to_owned(),
                    position: DiagramPosition { x: 245.0, y: 18.0 },
                },
            ],
            switches: HashMap::from([
                (
                    "SW_WEST".to_owned(),
                    DiagramSwitch {
                        switch_id: "SW_WEST".to_owned(),
                        position: DiagramPosition { x: -220.0, y: 0.0 },
                    },
                ),
                (
                    "SW_EAST".to_owned(),
                    DiagramSwitch {
                        switch_id: "SW_EAST".to_owned(),
                        position: DiagramPosition { x: 220.0, y: 0.0 },
                    },
                ),
            ]),
        },
    }
}
