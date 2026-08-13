<script lang="ts">
	import { Application, System } from '@wailsio/runtime';
	import { uiState } from '../state/ui.svelte';

	function onClickStartSimulation() {
		uiState.currentPage = 'selectSimulation';
	}

	function onClickLoadSimulation() {
		//uiState.currentPage = 'loadSimulation'
	}

	function onClickSettings() {
		uiState.currentPage = 'settings';
	}

	function onClickScenarioEditor() {
		uiState.currentPage = 'scenarioEditor';
	}

	function onClickExit() {
		Application.Quit();
	}
</script>

<div class="mx-20 mt-15">
	<div>
		<h1 class="text-2xl font-bold text-white">OpenRCS</h1>
		<h2 class="text-xl text-white">Open Source Rail Control Simulator</h2>
	</div>

	<div class="mt-20 flex max-w-xs flex-col gap-3">
		<button
			class="rounded bg-primary px-4 py-2 font-medium text-primary-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
			on:click={onClickStartSimulation}
		>
			Start a Simulation
		</button>
		<button
			class="rounded bg-primary px-4 py-2 font-medium text-primary-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
			on:click={onClickLoadSimulation}
		>
			Load a Simulation
		</button>
		<button
			class="rounded bg-primary px-4 py-2 font-medium text-primary-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
			on:click={onClickScenarioEditor}
		>
			Scenario Editor
		</button>
        <button
			class="rounded bg-primary px-4 py-2 font-medium text-primary-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
			on:click={onClickSettings}
		>
			Settings
		</button>

		<button
			class="rounded bg-primary px-4 py-2 font-medium text-primary-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
			on:click={onClickExit}
		>
			Exit
		</button>
	</div>
</div>
