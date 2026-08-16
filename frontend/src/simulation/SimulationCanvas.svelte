<script lang="ts">
	import type { TrackCircuit } from '../../bindings/github.com/chrisjm66/openrcs/internal/layout';
	import type { DiagramTrack } from '../../bindings/github.com/chrisjm66/openrcs/internal/signal_diagram';
	import { simulationState } from '../state/simulation.svelte';

	let diagram = simulationState.scenario?.SignallingDiagram;

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

			if (!diagram || !diagram.Tracks) {
				console.log('null');
				return;
			} else {
				Object.entries(diagram.Tracks).forEach(([id, track]) => {
					console.log(id);
					if (track && track.DiagramPositions) {
						console.log(track);

						for (let i = 1; i < track.DiagramPositions.length; i++) {
							const previousPosition = track.DiagramPositions[i - 1];
							const currentPosition = track.DiagramPositions[i];
							context.beginPath();
							context.moveTo(previousPosition.X, previousPosition.Y);
							context.lineTo(currentPosition.X, currentPosition.Y);
							context.lineWidth = 5;
							context.strokeStyle = 'darkgrey';
							context.stroke();
							// TODO perform state lookup to add headcode, color, etc.
						}
					} else {
						console.log('circuit is null');
					}
				});
			}
		});
	}
</script>

<canvas bind:this={canvas} class="h-full w-full bg-black"></canvas>
