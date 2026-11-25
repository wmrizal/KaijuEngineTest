package engine

import (
	"fmt"
	"kaijuengine/game"
	"kaijuengine/input"
	"kaijuengine/rendering"
)

// Engine represents the core game engine
type Engine struct {
	title      string
	width      int
	height     int
	running    bool
	deltaTime  float64
	targetFPS  int
	renderer   *rendering.Renderer
	inputMgr   *input.Manager
	gameState  *game.GameState
}

// NewEngine creates a new game engine instance
func NewEngine(title string, width, height int) *Engine {
	return &Engine{
		title:     title,
		width:     width,
		height:    height,
		running:   false,
		targetFPS: 60,
		renderer:  rendering.NewRenderer(title, width, height),
		inputMgr:  input.NewManager(),
		gameState: game.NewGameState(),
	}
}

// Run starts the game loop using Ebiten
func (e *Engine) Run() error {
	fmt.Printf("Starting %s (%dx%d) at %d FPS\n", e.title, e.width, e.height, e.targetFPS)
	
	e.running = true

	// Run the renderer (which uses Ebiten for the game loop)
	return e.renderer.Run()
}

// Update updates game logic
func (e *Engine) Update(deltaTime float64) {
	e.gameState.Update(deltaTime)
}

// Render renders a frame
func (e *Engine) Render() {
	e.renderer.Clear()
	e.renderer.Present()
}

// GetGameState returns the game state
func (e *Engine) GetGameState() *game.GameState {
	return e.gameState
}
