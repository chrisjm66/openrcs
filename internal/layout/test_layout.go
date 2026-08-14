package layout

func CreateTestLayout() *RailwayLayout {
	// This is a very simple layout. Imagine a simple junction, where one track splits into two. The mainline (tracks 1, 2) are protected by a single track circuit, while the other line also has its own track circuit. One signal protects the one circuit, and another signal protects the other. There is a junction which connects the two tracks with the common route being track 1, the normal route being track 2, and the reverse being track 3.

	tracks := map[TrackId]*Track{
		"1": {
			Length: 100,
		},
		"2": {
			Length: 200,
		},
		"3": {
			Length: 50,
		},
	}

	trackCircuits := map[TrackCircuitId]*TrackCircuit{
		"TC1": {
			StartPosition: TrackPosition{
				TrackId: "1",
				Offset:  0,
			},
			EndPosition: TrackPosition{
				TrackId: "2",
				Offset:  50,
			},
		},
		"TC2": {
			StartPosition: TrackPosition{
				TrackId: "3",
				Offset:  0,
			},
			EndPosition: TrackPosition{
				TrackId: "3",
				Offset:  50,
			},
		},
	}

	signals := map[SignalId]*Signal{
		"S001": {
			Protects: []TrackCircuitId{"1", "2"},
			Position: TrackPosition{
				TrackId: "1",
				Offset:  0,
			},
		},
	}

	switches := map[SwitchId]*Switch{
		"P001": {
			Common: TrackEndpoint{
				TrackId:  "1",
				Endpoint: End,
			},
			Normal: TrackEndpoint{
				TrackId:  "3",
				Endpoint: Start,
			},
			Reverse: TrackEndpoint{
				TrackId:  "2",
				Endpoint: Start,
			},
		},
	}

	layout := &RailwayLayout{
		Tracks:          tracks,
		TrackCircuits:   trackCircuits,
		Signals:         signals,
		Switches:        switches,
	}

	return layout
}
