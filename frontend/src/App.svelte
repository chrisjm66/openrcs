<script lang="ts">
	import './app.css';
	import Homepage from './home/Homepage.svelte';
	import { uiState } from './state/ui.svelte';

	let { currentPage } = $derived(uiState);
</script>

{#if currentPage == 'home'}
	<Homepage />
{:else if currentPage == 'settings'}
	<h1>settings</h1>
{:else if currentPage == 'editor'}
	<h1>editor</h1>
{:else if currentPage == 'simulation'}
	<h1>sim</h1>
{:else if currentPage == 'selectSimulation'}
	<h1>balls</h1>
{/if}
