<script lang="ts">
	import { uiState } from '../state/ui.svelte';
	import { SimulationService } from '../../bindings/github.com/chrisjm66/openrcs/internal/simulation';
	import { Call } from '@wailsio/runtime';
	import { loadSimulation } from '../state/simulation.svelte';
	import type {
		Scenario,
		ScenarioId
	} from '../../bindings/github.com/chrisjm66/openrcs/internal/simulation';

	let selectableScenarios = $state<Record<ScenarioId, Scenario | undefined>>();
	let selectedSimulationId = $state<ScenarioId>('');

	async function loadSimulations() {
		selectableScenarios = (await SimulationService.GetScenarios()) ?? {};
	}

	loadSimulations();

	function onClickReturnToMenu() {
		uiState.currentPage = 'home';
	}

	function onClickSelectSimulation(simulationId: string) {
		selectedSimulationId = simulationId;
	}

	async function onClickStartSimulation() {
		try {
			await loadSimulation(selectedSimulationId as string);
			uiState.currentPage = 'simulation';
		} catch (error: unknown) {
			if (error instanceof Call.RuntimeError) {
				console.log(('Error: simulation not found: ' + selectedSimulationId) as string);
			}
		}
	}
</script>

<div class="flex h-full w-full flex-row gap-y-5">
	<!-- Left panel -->
	<div class="flex h-full w-1/4 flex-col">
		<button
			onclick={onClickReturnToMenu}
			class="w-max rounded bg-primary p-1 text-primary-foreground transition-colors hover:bg-accent"
			>Return to Menu</button
		>
		<h1 class="my-5 text-2xl">Select a Simulation</h1>

		<div class="flex max-h-2/3 w-50 flex-col gap-y-2">
			{#each Object.entries(selectableScenarios ?? {}) as [id, scenario] (id)}
				{#if scenario}
					<button
						class="w-full rounded bg-primary p-2 text-primary-foreground transition-colors hover:bg-accent"
						onclick={() => onClickSelectSimulation(id)}
						value={id}
					>
						{scenario.name}
					</button>
				{/if}
			{/each}
		</div>
	</div>

	{#if selectedSimulationId != ''}
		<div class="h-full w-1/2">
			{#if selectableScenarios && selectableScenarios[selectedSimulationId]}
				<h2 class="text-xl">{selectableScenarios[selectedSimulationId]?.name}</h2>
				<p>{selectableScenarios[selectedSimulationId]?.description}</p>

				<button
					onclick={() => onClickStartSimulation()}
					class="transiton-colors absolute right-10 bottom-10 rounded bg-primary p-2 text-primary-foreground"
					>Start Simulation</button
				>
			{/if}
		</div>
	{/if}
</div>
