<script lang="ts">
	import { uiState } from '../state/ui.svelte';
	import { Call } from '@wailsio/runtime';
	import { loadSimulation } from '../state/simulation.svelte';
	import { SimulationService, type Scenario } from '../../bindings/github.com/chrisjm66/openrcs/internal/simulation';

	let selectableScenarios = $state<Scenario[]>();
	let selectedSimulationIndex = $state<number | undefined>(undefined);

	loadSimulations()

	async function loadSimulations() {
		const scenarios = (await SimulationService.GetScenarios()) ?? {};
		console.log(scenarios)
		if (scenarios != null && scenarios[0] != null) {
			selectableScenarios = scenarios[0]
		}	
	}

	function onClickReturnToMenu() {
		uiState.currentPage = 'home';
	}

	function onClickSelectSimulation(index: string) {
		selectedSimulationIndex = Number(index);
	}

	async function onClickStartSimulation(scenario: Scenario) {
		try {
			await loadSimulation(scenario.id);
			uiState.currentPage = 'simulation';
		} catch (error: unknown) {
			if (error instanceof Call.RuntimeError) {
				console.log(('Error: simulation not found: ' + scenario.id) as string);
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
			{#each Object.entries(selectableScenarios ?? {}) as [index, scenario] (scenario.id)}
				{#if scenario}
					<button
						class="w-full rounded bg-primary p-2 text-primary-foreground transition-colors hover:bg-accent"
						onclick={() => onClickSelectSimulation(index)}
						value={index}
					>
						{scenario.name}
					</button>
				{/if}
			{/each}
		</div>
	</div>

	{#if selectedSimulationIndex !== undefined}
		<div class="h-full w-1/2">
			{#if selectableScenarios && selectedSimulationIndex !== undefined && selectableScenarios[selectedSimulationIndex]}
				{@const selectedScenario = selectableScenarios[selectedSimulationIndex]}
				<h2 class="text-xl">{selectedScenario?.name}</h2>
				<p>{selectedScenario?.description}</p>

				<button
					onclick={() => onClickStartSimulation(selectedScenario)}
					class="transiton-colors absolute right-10 bottom-10 rounded bg-primary p-2 text-primary-foreground"
					>Start Simulation</button
				>
			{/if}
		</div>
	{/if}
</div>
