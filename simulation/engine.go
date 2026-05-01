package simulation

import (
	"fmt"
	"time"
)

func RunSimulation() {
	go initializeTickLoop()
}

func initializeTickLoop() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		time.Sleep(50 * time.Millisecond)
		fmt.Print("simloop running")
	}
}

func stepSimulation(queuedCommands []Command, worldState WorldState) {
	for _, command := range queuedCommands {
		command.Apply(&worldState)
	}
}
