package rendering

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Renderer handles all rendering operations using Ebiten
type Renderer struct {
	title  string
	width  int
	height int
	game   *Game
}

// Game implements ebiten.Game interface
type Game struct {
	renderer *Renderer
	frameNum int
}

// NewRenderer creates a new renderer
func NewRenderer(title string, width, height int) *Renderer {
	r := &Renderer{
		title:  title,
		width:  width,
		height: height,
	}
	r.game = &Game{renderer: r}
	return r
}

// Update updates game state (required by ebiten.Game)
func (g *Game) Update() error {
	g.frameNum++
	return nil
}

// Draw renders the frame (required by ebiten.Game)
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear background
	screen.Fill(color.RGBA{10, 10, 20, 255})

	// Draw title bar
	g.renderer.DrawFilledRect(screen, 0, 0, g.renderer.width, 50, color.RGBA{30, 30, 60, 255})
	g.renderer.DrawText(screen, "🦖 KaijuEngine Demo - Now with Rendering!", 20, 15)

	// Draw FPS
	g.renderer.DrawText(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()), 20, 300)

	// Draw animated shapes (bouncing circles)
	y := 100
	for i := 0; i < 5; i++ {
		x := 50 + i*150
		offset := (g.frameNum + i*20) % 100
		g.renderer.DrawFilledCircle(screen, float64(x), float64(y+offset), 25, color.RGBA{uint8(50 + i*40), uint8(100 + i*20), 200, 255})
	}

	// Draw a rotating square in center
	centerX := float64(g.renderer.width / 2)
	centerY := float64(g.renderer.height / 2)
	g.renderer.DrawFilledRect(screen, int(centerX)-25, int(centerY)-25, 50, 50, color.RGBA{255, 100, 50, 255})

	// Draw instructions
	g.renderer.DrawText(screen, "Press ESC or close window to exit", 20, g.renderer.height-30)
}

// Layout returns the game's logical screen size (required by ebiten.Game)
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.renderer.width, g.renderer.height
}

// Run starts the renderer
func (r *Renderer) Run() error {
	ebiten.SetWindowSize(r.width, r.height)
	ebiten.SetWindowTitle(r.title)
	return ebiten.RunGame(r.game)
}

// Clear clears the screen
func (r *Renderer) Clear() {
	// Ebiten handles clearing in Draw method
}

// DrawText draws text at the given position
func (r *Renderer) DrawText(screen *ebiten.Image, text string, x, y int) {
	// Use Ebiten's debug text rendering for simplicity
	ebitenutil.DebugPrintAt(screen, text, x, y)
}

// DrawFilledRect draws a filled rectangle
func (r *Renderer) DrawFilledRect(screen *ebiten.Image, x, y, width, height int, col color.Color) {
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(width), float32(height), col, true)
}

// DrawFilledCircle draws a filled circle
func (r *Renderer) DrawFilledCircle(screen *ebiten.Image, x, y, radius float64, col color.Color) {
	vector.DrawFilledCircle(screen, float32(x), float32(y), float32(radius), col, true)
}

// Present presents the frame
func (r *Renderer) Present() {
	// Ebiten handles presentation automatically
}

// Close cleans up renderer resources
func (r *Renderer) Close() error {
	fmt.Println("Renderer closed")
	return nil
}
