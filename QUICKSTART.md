# KaijuEngine - Quick Start Guide

## 🦖 Welcome to KaijuEngine!

A lightweight, modular 2D game engine written in Go.

## Quick Start (5 minutes)

### 1. Build the Engine
```bash
cd D:\repos\KaijuEngine
go build -o KaijuEngine.exe
```

### 2. Run the Engine
```bash
.\KaijuEngine.exe
```

You should see:
```
🦖 KaijuEngine - Game Engine
============================
Starting KaijuEngine (800x600) at 60 FPS
Engine shutdown complete
```

## Project Structure

```
KaijuEngine/
├── main.go              ← Entry point - starts the engine
├── go.mod               ← Go module manifest
├── engine/              ← Core game loop
│   └── engine.go        ← Engine implementation
├── game/                ← Game state & entities
│   └── state.go         ← Entity and GameState types
├── rendering/           ← Graphics system (placeholder)
│   └── renderer.go      ← Renderer interface
└── input/               ← Input management
    └── manager.go       ← Input handler
```

## Key Components

### Engine (`engine/engine.go`)
Manages the game loop:
- Fixed 60 FPS update rate
- Automatic delta time calculation
- Coordinates Update → Render cycle

### GameState (`game/state.go`)
Holds game entities and logic:
- Entity management
- Position/velocity updates
- Easy to extend with custom logic

### Renderer (`rendering/renderer.go`)
Graphics interface (ready for implementation):
- Currently placeholders
- Ready for graphics library integration (Ebiten, raylib-go, etc.)

### InputManager (`input/manager.go`)
Input handling interface:
- Extensible design
- Ready for keyboard/mouse/gamepad support

## Common First Steps

### Add an Entity
```go
entity := &game.Entity{
    ID:     1,
    X:      100,
    Y:      100,
    Width:  32,
    Height: 32,
    Active: true,
    VelX:   100,  // pixels/sec
    VelY:   0,
}
gameState.AddEntity(entity)
```

### Extend the Engine
Edit `engine/engine.go` → `Update()` method to add custom game logic.

### Add Graphics
1. Choose a library: **Ebiten** recommended
2. Install: `go get github.com/hajimehoshi/ebiten/v2`
3. Update `rendering/renderer.go` with actual graphics calls
4. Update game loop in `engine/engine.go` to use renderer

## Next: Building Your First Game

See [DEVGUIDE.md](DEVGUIDE.md) for:
- Detailed architecture overview
- Development roadmap
- Graphics library recommendations
- Example implementations

## Useful Commands

```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...

# Run tests (when added)
go test ./...

# Clean build
go clean
go build -o KaijuEngine.exe

# Check module dependencies
go list -m all
```

## Recommended Next Features

1. **Graphics Library Integration** - Render sprites and text
2. **Keyboard Input** - Handle player controls
3. **Collision Detection** - Basic AABB collision
4. **Sound System** - Play effects and music
5. **Tile Maps** - Build level layouts

## Resources

- [Go Official Docs](https://golang.org/doc/)
- [Ebiten Documentation](https://ebiten.org/)
- [Game Programming Patterns](https://gameprogrammingpatterns.com/)

## License

MIT

---

**Happy game development! 🎮**
