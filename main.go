package main

import (
	"fmt"
	"kaijuengine/engine"
)

func main() {
	fmt.Println("🦖 KaijuEngine - Game Engine")
	fmt.Println("============================")

	// Initialize the engine
	gameEngine := engine.NewEngine("KaijuEngine", 800, 600)
	
	// Run the game loop
	if err := gameEngine.Run(); err != nil {
		fmt.Printf("Engine error: %v\n", err)
	}
}
