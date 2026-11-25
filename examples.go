package main

import (
	"fmt"
	"kaijuengine/engine"
	"kaijuengine/game"
)

// This file demonstrates how to extend the engine with custom game logic.
// It's not used in the main build, but serves as a reference.

// ExampleGame shows how to create a simple game using KaijuEngine
func ExampleGameUsage() {
	// Create engine instance
	gameEngine := engine.NewEngine("KaijuEngine Demo", 800, 600)
	
	// Access game state
	gameState := gameEngine.GetGameState() // This would need to be added to engine
	
	// Create a player entity
	player := &game.Entity{
		ID:     1,
		X:      400,
		Y:      300,
		Width:  32,
		Height: 32,
		Active: true,
		VelX:   0,
		VelY:   0,
	}
	
	// Add player to game state
	gameState.AddEntity(player)
	
	// Create some enemy entities
	for i := 0; i < 5; i++ {
		enemy := &game.Entity{
			ID:     i + 2,
			X:      float64(100 + i*100),
			Y:      100,
			Width:  32,
			Height: 32,
			Active: true,
			VelX:   50,
			VelY:   0,
		}
		gameState.AddEntity(enemy)
	}
	
	// Run the game
	if err := gameEngine.Run(); err != nil {
		fmt.Printf("Game error: %v\n", err)
	}
}

// HandleInput demonstrates how to handle player input
func HandleInput(player *game.Entity) {
	// This would be called in the engine's update loop
	// In a real implementation, check actual keyboard state
	
	const moveSpeed = 300.0 // pixels per second
	
	// Example: Set velocity based on input
	// if input.IsKeyPressed(input.KEY_LEFT) {
	//     player.VelX = -moveSpeed
	// }
	// if input.IsKeyPressed(input.KEY_RIGHT) {
	//     player.VelX = moveSpeed
	// }
}

// HandleCollision demonstrates basic collision detection
func HandleCollision(entity1, entity2 *game.Entity) bool {
	// Simple AABB (Axis-Aligned Bounding Box) collision
	return entity1.X < entity2.X+entity2.Width &&
		entity1.X+entity1.Width > entity2.X &&
		entity1.Y < entity2.Y+entity2.Height &&
		entity1.Y+entity1.Height > entity2.Y
}

// UpdateGameLogic demonstrates how to extend the update loop
func UpdateGameLogic(gameState *game.GameState, deltaTime float64) {
	entities := gameState.GetEntities()
	
	// Update each entity
	for i, entity := range entities {
		if !entity.Active {
			continue
		}
		
		// Update position
		entity.X += entity.VelX * deltaTime
		entity.Y += entity.VelY * deltaTime
		
		// Check bounds
		if entity.X < 0 || entity.X > 800 {
			entity.VelX *= -1 // Bounce off walls
		}
		if entity.Y < 0 || entity.Y > 600 {
			entity.VelY *= -1 // Bounce off walls
		}
		
		// Check collisions with other entities
		for j := i + 1; j < len(entities); j++ {
			other := entities[j]
			if other.Active && HandleCollision(entity, other) {
				// Handle collision
				fmt.Printf("Collision between entity %d and %d\n", entity.ID, other.ID)
				// Implement your collision response logic here
			}
		}
	}
}
