<script lang="ts">
	import { uiState } from "../state/ui.svelte";

  function onClickReturnToMenu() {
    uiState.currentPage = 'home'
  }
</script>

<div>
  <button on:click={onClickReturnToMenu}>Return to Menu</button>
  <h1>Settings</h1>

  <input type="text" placeholder="probably a dropdown idk">
</div>
