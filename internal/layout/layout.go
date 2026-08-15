package layout

import (
	"errors"
)

func GetAvailableLayouts() []RailwayLayoutId {
	layouts := []RailwayLayoutId{}

	for id := range railwayLayouts() {
		layouts = append(layouts, id)
	}

	return layouts
}

func GetLayout(railwayLayoutId RailwayLayoutId) (RailwayLayout, error) {
	layout, ok := railwayLayouts()[railwayLayoutId]

	if !ok {
		return RailwayLayout{}, errors.New("Simulation " + string(railwayLayoutId) + " not found")
	}

	return *layout, nil
}

func railwayLayouts() map[RailwayLayoutId]*RailwayLayout {
	return map[RailwayLayoutId]*RailwayLayout{
		"Test Layout 1": CreateTestLayout(),
		"Test Layout 2": CreateTestLayout(),
	}
}

type RailwayLayoutId string

type RailwayLayout struct {
	Tracks        map[TrackId]*Track
	Signals       map[SignalId]*Signal
	TrackCircuits map[TrackCircuitId]*TrackCircuit
	Switches      map[SwitchId]*Switch
}
