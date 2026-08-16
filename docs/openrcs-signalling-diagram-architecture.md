# OpenRCS signalling diagram architecture

```mermaid
flowchart TB
    subgraph Scenario["Static scenario definition"]
        Layout["Railway layout\nPhysical circuits, signals, switches, and topology"]
        Interlocking["Interlocking rules\nRoutes, point requirements, and exclusive resources"]
        Diagram["Signalling diagram\nDisplay segments, symbols, labels, and berths"]
    end

    subgraph Runtime["Simulation runtime"]
        State["World state\nOccupancy, point positions, route locks, signal aspects, and trains"]
        Engine["Simulation engine\nProcesses requests and evaluates rules"]
        Panel["Operator panel\nRenders the schematic from live state"]
    end

    Layout -->|"layout context"| Engine
    Interlocking -->|"route rules"| Engine
    Diagram -->|"static appearance"| Panel
    Panel -->|"route request"| Engine
    Engine -->|"state updates"| State
    State -->|"state snapshot"| Panel

    Note["Display rule: a segment normally references one circuit.\nIt may reference several only when the panel should show them as one operational section."]
    Diagram -.-> Note
```

The signalling diagram is static scenario data. The operator panel derives colours and labels from `WorldState`; it does not make route or safety decisions.
