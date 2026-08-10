<script lang="ts">
	import './app.css';
	let test = $state<string>('');
</script>

<div>
	<div class="mt-15">
		<h1 class="text-2xl font-bold text-white">OpenRCS</h1>
		<h2 class="text-xl text-white">Open Source Rail Control Simulator</h2>
	</div>

	<div class="mt-20 flex max-w-xs flex-col gap-3">
		<button
			class="rounded bg-primary px-4 py-2 font-medium text-primary-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
		>
			Start a Simulation
		</button>
		<button
			class="rounded bg-secondary px-4 py-2 font-medium text-secondary-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
		>
			Load a Simulation
		</button>
		<button
			class="rounded border border-border bg-muted px-4 py-2 font-medium text-muted-foreground transition-colors hover:text-foreground"
		>
			Scenario Editor
		</button>
		<button
			class="rounded border border-border bg-muted px-4 py-2 font-medium text-muted-foreground transition-colors hover:text-foreground"
		>
			Settings
		</button>
	</div>
</div>
