<script lang="ts">
	import './app.css';
	import HomeContainer from './home/HomeContainer.svelte';
	import Homepage from './home/Homepage.svelte';
	import SelectSimulation from './home/SelectSimulation.svelte';
	import Simulation from './simulation/Simulation.svelte';
	import { uiState } from './state/ui.svelte';

	let { currentPage } = $derived(uiState);
</script>

<div class="h-full w-full">
	{#if currentPage == 'home'}
		<HomeContainer>
			<Homepage />
		</HomeContainer>
	{:else if currentPage == 'settings'}
		<HomeContainer>
			<h1>settings here</h1>
		</HomeContainer>
	{:else if currentPage == 'editor'}
		<h1>editor</h1>
	{:else if currentPage == 'simulation'}
		<Simulation />
	{:else if currentPage == 'selectSimulation'}
		<HomeContainer>
			<SelectSimulation />
		</HomeContainer>
	{:else}
		<h1>no page</h1>
	{/if}
</div>
