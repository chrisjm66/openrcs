import type { WorldState } from '../../bindings/github.com/chrisjm66/openrcs/internal/state';
import {
	SimulationService,
	type Scenario,
	type ScenarioId
} from '../../bindings/github.com/chrisjm66/openrcs/internal/simulation';

export const simulationState = $state<SimulationState>({
	worldState: undefined,
	scenario: undefined
});

export async function loadSimulation(scenarioId: ScenarioId) {
	const error = await SimulationService.NewRuntime(scenarioId as string);

	if (error != null) {
		console.log('Error: ' + error.message);
	}

	simulationState.scenario = await SimulationService.GetCurrentScenario();
	console.log(simulationState.scenario?.layout);
}

interface SimulationState {
	worldState: WorldState | undefined;
	scenario: Scenario | undefined;
}
