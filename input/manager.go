package input

import "github.com/hajimehoshi/ebiten/v2"

// Manager manages input events
type Manager struct {
	quit           bool
	horizontalAxis float64
}

// NewManager creates a new input manager
func NewManager() *Manager {
	return &Manager{}
}

// Update refreshes the current keyboard state.
func (m *Manager) Update() {
	m.update(
		ebiten.IsKeyPressed(ebiten.KeyEscape),
		ebiten.IsKeyPressed(ebiten.KeyA),
		ebiten.IsKeyPressed(ebiten.KeyD),
	)
}

func (m *Manager) update(escape, left, right bool) {
	if escape {
		m.quit = true
	}

	m.horizontalAxis = 0
	if left {
		m.horizontalAxis--
	}
	if right {
		m.horizontalAxis++
	}
}

// GetQuit checks if quit signal was received
func (m *Manager) GetQuit() bool {
	return m.quit
}

// GetHorizontalAxis returns -1 for left, 1 for right, and 0 when idle.
func (m *Manager) GetHorizontalAxis() float64 {
	return m.horizontalAxis
}

// SetQuit sets the quit flag
func (m *Manager) SetQuit(quit bool) {
	m.quit = quit
}
