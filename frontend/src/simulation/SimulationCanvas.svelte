<script lang="ts">
	import type { TrackCircuit } from '../../bindings/github.com/chrisjm66/openrcs/internal/layout';
	import { simulationState } from '../state/simulation.svelte';

	let layout = simulationState.layout;

	let canvas: HTMLCanvasElement;

	initializeCanvas();
	subscribeToLayoutUpdates();

	function initializeCanvas() {
		$effect(() => {
			const context = canvas.getContext('2d');

			const rect = canvas.getBoundingClientRect();
			const dpr = window.devicePixelRatio || 1;

			canvas.width = rect.width * dpr;
			canvas.height = rect.height * dpr;

			context?.scale(dpr, dpr);
		});
	}
	function subscribeToLayoutUpdates() {
		$effect(() => {
			const context = canvas.getContext('2d');

			if (context == null) {
				return;
			}

			if (!layout || !layout.TrackCircuits) {
				console.log('null');
				return;
			} else {
				Object.entries(layout.TrackCircuits).forEach(([id, trackCircuit]) => {
					const circuit = trackCircuit as TrackCircuit;
					console.log(id);
					if (circuit) {
						console.log(circuit);
						context.beginPath();
						context.moveTo(100, 300);
						context.lineTo(100, 300 + circuit.EndPosition.Offset);
						context.lineWidth = 5;
						context.strokeStyle = 'white';
						context.stroke();

						context.beginPath();
						context.fillStyle = 'red';
						context.fillText(id, 400, 300 + circuit.EndPosition.Offset / 2);
						context.stroke();
					} else {
						console.log('circuit is null');
					}
				});
			}
		});
	}
</script>

<canvas bind:this={canvas} class="h-full w-full bg-black"></canvas>
