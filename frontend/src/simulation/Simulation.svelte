<script lang="ts">
	import { simulationState } from '../state/simulation.svelte';
	import Footer from './Footer.svelte';
	import SimulationCanvas from './SimulationCanvas.svelte';
	import TopBar from './TopBar.svelte';
</script>

<div>
	<!-- main container -->
	<div class="absolute h-screen w-screen bg-black">
		<div class="top-0">
			<TopBar />
		</div>

		<div class="w-full h-full">
			<SimulationCanvas diagram={simulationState.scenario?.signallingDiagram} />
		</div>

		<div class="absolute bottom-0">
			<Footer />
		</div>
	</div>
</div>
