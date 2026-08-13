package layout

type RailwayLayout struct {
	Tracks        map[TrackId]*Track
	Signals       map[SignalId]*Signal
	TrackCircuits map[TrackCircuitId]*TrackCircuit
	Switches      map[SwitchId]*Switch
}
