package input

import "testing"

func TestHorizontalAxis(t *testing.T) {
	tests := []struct {
		name        string
		left, right bool
		want        float64
	}{
		{name: "idle", want: 0},
		{name: "move left", left: true, want: -1},
		{name: "move right", right: true, want: 1},
		{name: "opposing keys cancel", left: true, right: true, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewManager()
			manager.update(false, tt.left, tt.right)

			if got := manager.GetHorizontalAxis(); got != tt.want {
				t.Fatalf("GetHorizontalAxis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEscapeRequestsQuit(t *testing.T) {
	manager := NewManager()
	manager.update(true, false, false)

	if !manager.GetQuit() {
		t.Fatal("GetQuit() = false after Escape was pressed")
	}
}
