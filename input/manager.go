package input

// Manager manages input events
type Manager struct {
	quit bool
}

// NewManager creates a new input manager
func NewManager() *Manager {
	return &Manager{
		quit: false,
	}
}

// GetQuit checks if quit signal was received
func (m *Manager) GetQuit() bool {
	// In a real implementation, this would check for window close events
	// and keyboard input (e.g., ESC key)
	return m.quit
}

// SetQuit sets the quit flag
func (m *Manager) SetQuit(quit bool) {
	m.quit = quit
}
