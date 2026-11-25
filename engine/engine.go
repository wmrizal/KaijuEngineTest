package engine

import (
	"fmt"
	"kaijuengine/game"
	"kaijuengine/input"
	"kaijuengine/rendering"
	"time"
)

// Engine represents the core game engine
type Engine struct {
	title      string
	width      int
	height     int
	running    bool
	deltaTime  float64
	targetFPS  int
	lastFrame  time.Time
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

// Run starts the game loop
func (e *Engine) Run() error {
	fmt.Printf("Starting %s (%dx%d) at %d FPS\n", e.title, e.width, e.height, e.targetFPS)
	
	e.running = true
	e.lastFrame = time.Now()
	
	frameTargetDuration := time.Duration(1000/e.targetFPS) * time.Millisecond

	// Game loop
	for e.running && e.frameCount() < 300 { // Run for ~5 seconds for demo
		frameStart := time.Now()

		// Handle input
		if e.inputMgr.GetQuit() {
			e.running = false
		}

		// Update game state
		e.Update(e.deltaTime)

		// Render frame
		e.Render()

		// Frame rate limiting
		elapsedTime := time.Since(frameStart)
		if elapsedTime < frameTargetDuration {
			time.Sleep(frameTargetDuration - elapsedTime)
		}

		// Calculate delta time
		now := time.Now()
		e.deltaTime = now.Sub(e.lastFrame).Seconds()
		e.lastFrame = now
	}

	fmt.Println("Engine shutdown complete")
	return nil
}

// Update updates game logic
func (e *Engine) Update(deltaTime float64) {
	e.gameState.Update(deltaTime)
}

// Render renders a frame
func (e *Engine) Render() {
	e.renderer.Clear()
	e.renderer.DrawText(fmt.Sprintf("FPS: %.0f", 1.0/e.deltaTime), 10, 10)
	e.renderer.DrawText("KaijuEngine - Basic Game Engine", 10, 30)
	e.renderer.Present()
}

// frameCount is a simple frame counter for demo purposes
func (e *Engine) frameCount() int {
	return int(e.lastFrame.Unix() * 100 % 300)
}
