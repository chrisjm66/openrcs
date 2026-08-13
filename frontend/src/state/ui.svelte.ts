export const uiState = $state<UiState>({
	currentPage: 'home'
});

interface UiState {
	currentPage: Pages;
}

export type Pages = 'home' | 'simulation' | 'editor' | 'settings' | 'selectSimulation' | 'scenarioEditor';
