package engine

import (
	"fmt"
	"image/color"
	"kaijuengine/game"
	"kaijuengine/input"
	"kaijuengine/rendering"

	"github.com/hajimehoshi/ebiten/v2"
)

// Engine represents the core game engine
type Engine struct {
	title     string
	width     int
	height    int
	running   bool
	deltaTime float64
	targetFPS int
	renderer  *rendering.Renderer
	inputMgr  *input.Manager
	gameState *game.GameState
	player    *game.Entity
}

// NewEngine creates a new game engine instance
func NewEngine(title string, width, height int) *Engine {
	e := &Engine{
		title:     title,
		width:     width,
		height:    height,
		running:   false,
		targetFPS: 60,
		renderer:  rendering.NewRenderer(title, width, height),
		inputMgr:  input.NewManager(),
		gameState: game.NewGameState(),
	}
	e.player = &game.Entity{
		ID:     1,
		X:      float64(width/2 - 25),
		Y:      float64(height/2 - 25),
		Width:  50,
		Height: 50,
		Active: true,
	}
	e.gameState.AddEntity(e.player)
	e.renderer.SetUpdateFunc(e.update)
	e.renderer.SetDrawFunc(e.draw)
	return e
}

// Run starts the game loop using Ebiten
func (e *Engine) Run() error {
	fmt.Printf("Starting %s (%dx%d) at %d FPS\n", e.title, e.width, e.height, e.targetFPS)

	e.running = true
	defer func() {
		e.running = false
	}()

	// Run the renderer (which uses Ebiten for the game loop)
	return e.renderer.Run()
}

func (e *Engine) update(deltaTime float64) error {
	const movementSpeed = 300.0

	e.inputMgr.Update()
	if e.inputMgr.GetQuit() {
		return ebiten.Termination
	}

	e.player.VelX = e.inputMgr.GetHorizontalAxis() * movementSpeed
	e.Update(deltaTime)

	maxX := float64(e.width) - e.player.Width
	if e.player.X < 0 {
		e.player.X = 0
	} else if e.player.X > maxX {
		e.player.X = maxX
	}

	return nil
}

// Update updates game logic
func (e *Engine) Update(deltaTime float64) {
	e.gameState.Update(deltaTime)
}

func (e *Engine) draw(screen *ebiten.Image, frameNum int) {
	screen.Fill(color.RGBA{10, 10, 20, 255})

	e.renderer.DrawFilledRect(screen, 0, 0, e.width, 50, color.RGBA{30, 30, 60, 255})
	e.renderer.DrawText(screen, "KaijuEngine Demo - Keyboard Input", 20, 15)
	e.renderer.DrawText(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()), 20, 300)

	y := 100
	for i := 0; i < 5; i++ {
		x := 50 + i*150
		offset := (frameNum + i*20) % 100
		e.renderer.DrawFilledCircle(screen, float64(x), float64(y+offset), 25, color.RGBA{uint8(50 + i*40), uint8(100 + i*20), 200, 255})
	}

	for _, entity := range e.gameState.GetEntities() {
		if entity.Active {
			e.renderer.DrawFilledRect(screen, int(entity.X), int(entity.Y), int(entity.Width), int(entity.Height), color.RGBA{255, 100, 50, 255})
		}
	}

	e.renderer.DrawText(screen, "A/D: move left/right    ESC: exit", 20, e.height-30)
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
