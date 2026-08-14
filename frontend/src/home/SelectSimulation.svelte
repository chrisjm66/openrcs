<script lang="ts">
	import { uiState } from '../state/ui.svelte';
	import { SimulationService } from '../../bindings/github.com/chrisjm66/openrcs/internal/simulation';

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
</script>

<div class="flex flex-row h-full w-full gap-y-5">
	<!-- Left panel -->
	<div class="flex flex-col h-full w-1/4">
		<button
			onclick={onClickReturnToMenu}
			class="w-max rounded bg-primary p-1 text-primary-foreground transition-colors hover:bg-accent"
			>Return to Menu</button
		>
		<h1 class="text-2xl my-5">Select a Simulation</h1>

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
		<h2 class='text-xl'>{selectedSimulation}</h2>
    <p>Description here</p>

    <button class="rounded absolute bottom-10 right-10 bg-primary text-primary-foreground p-2">Start Simulation</button>
	</div>
  {/if}
</div>
