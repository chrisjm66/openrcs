<script lang="ts">
	import { uiState } from '../state/ui.svelte';
	import { SimulationService } from '../../bindings/github.com/chrisjm66/openrcs/internal/simulation';
	import { Call } from '@wailsio/runtime';
	import { loadSimulation } from '../state/simulation.svelte';

	let selectableSimulations = $state<string[]>([]);
	let selectedSimulation = $state<string>('');

	async function loadSimulations() {
		selectableSimulations = (await SimulationService.GetLayouts()) ?? [];
	}

	loadSimulations();

	function onClickReturnToMenu() {
		uiState.currentPage = 'home';
	}

	function onClickSelectSimulation(simulation: string) {
		selectedSimulation = simulation;
	}

	async function onClickStartSimulation() {
		try {
			await loadSimulation(selectedSimulation);
			uiState.currentPage = 'simulation';
		} catch (error: unknown) {
			if (error instanceof Call.RuntimeError) {
				console.log('Error: simulation not found: ' + selectedSimulation);
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
			{#each selectableSimulations as simulation (simulation)}
				<button
					class="w-full rounded bg-primary p-2 text-primary-foreground transition-colors hover:bg-accent"
					onclick={() => onClickSelectSimulation(simulation)}
					value={simulation}
				>
					{simulation}
				</button>
			{/each}
		</div>
	</div>

	{#if selectedSimulation != ''}
		<div class="h-full w-1/2">
			<h2 class="text-xl">{selectedSimulation}</h2>
			<p>Description here</p>

			<button
				onclick={() => onClickStartSimulation()}
				class="transiton-colors absolute right-10 bottom-10 rounded bg-primary p-2 text-primary-foreground"
				>Start Simulation</button
			>
		</div>
	{/if}
</div>
