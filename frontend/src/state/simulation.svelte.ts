import type {
	RailwayLayout,
	RailwayLayoutId
} from '../../bindings/github.com/chrisjm66/openrcs/internal/layout';
import type { WorldState } from '../../bindings/github.com/chrisjm66/openrcs/internal/state';
import { SimulationService } from '../../bindings/github.com/chrisjm66/openrcs/internal/simulation';

export const simulationState = $state<SimulationState>({
	worldState: null,
	layout: null
});

export async function loadSimulation(railwayLayoutId: RailwayLayoutId) {
	await SimulationService.NewRuntime(railwayLayoutId);
	simulationState.layout = await SimulationService.GetLayout();
	console.log(simulationState.layout);
}

interface SimulationState {
	worldState: WorldState | null;
	layout: RailwayLayout | null;
}
