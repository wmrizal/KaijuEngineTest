# KaijuEngine Development Guide

## Project Overview

**KaijuEngine** is a basic 2D game engine written in Go, designed to provide a foundation for building simple games with a focus on modularity and extensibility.

## Architecture

The engine follows a modular architecture divided into key packages:

### Core Packages

#### 1. **engine/** - Core Engine Loop
- `engine.go`: Main game engine with:
  - Game loop implementation
  - Frame rate management (60 FPS default)
  - Delta time calculation
  - Update/Render cycle

**Key Types:**
- `Engine`: Main engine struct that orchestrates the game loop

**Usage:**
```go
gameEngine := engine.NewEngine("KaijuEngine", 800, 600)
gameEngine.Run()
```

#### 2. **game/** - Game State & Entities
- `state.go`: Game state management with:
  - Entity management system
  - Game state updates
  - Basic entity lifecycle

**Key Types:**
- `Entity`: Represents game objects with position, velocity, and dimensions
- `GameState`: Container for all game entities and logic

**Usage:**
```go
entity := &game.Entity{
    ID: 1,
    X: 100,
    Y: 100,
    Width: 32,
    Height: 32,
    Active: true,
    VelX: 100,
    VelY: 0,
}
gameState.AddEntity(entity)
```

#### 3. **rendering/** - Graphics System
- `renderer.go`: Rendering interface with:
  - Screen clearing
  - Text drawing
  - Frame presentation

**Key Types:**
- `Renderer`: Handles all rendering operations

**Current State:** Placeholder implementation ready for graphics library integration

#### 4. **input/** - Input Handling
- `manager.go`: Input management with:
  - Quit signal handling
  - Extensible for keyboard/mouse input

**Key Types:**
- `Manager`: Manages input events

**Future:** Add keyboard, mouse, and gamepad support

## Development Roadmap

### Phase 1: Foundation ✅
- [x] Basic project structure
- [x] Game loop implementation
- [x] Entity management system

### Phase 2: Rendering (In Progress)
- [ ] Integrate graphics library (e.g., raylib-go, Ebiten, or SDL2)
- [ ] Implement sprite rendering
- [ ] Add camera system
- [ ] Support for tile maps

### Phase 3: Input System
- [ ] Keyboard input handling
- [ ] Mouse input handling
- [ ] Gamepad support
- [ ] Input mapping system

### Phase 4: Physics & Collision
- [ ] Collision detection
- [ ] Rigid body physics
- [ ] Gravity simulation

### Phase 5: Audio System
- [ ] Sound effects
- [ ] Background music
- [ ] Audio mixer

### Phase 6: Asset Management
- [ ] Asset loading system
- [ ] Resource caching
- [ ] Asset hot-reloading

## Building and Running

### Build the Engine
```bash
go build -o KaijuEngine.exe
```

### Run the Engine
```bash
./KaijuEngine.exe
```

### Run Tests (Future)
```bash
go test ./...
```

## Graphics Library Options

For the rendering system, consider these Go graphics libraries:

1. **Ebiten** - Most suitable for 2D games
   - Easy to use API
   - Good performance
   - Cross-platform
   - Install: `go get github.com/hajimehoshi/ebiten/v2`

2. **raylib-go** - Bindings for raylib
   - Lightweight
   - Good for 2D and 3D
   - Install: `go get github.com/gen2brain/raylib-go/raylib`

3. **SDL2** - Low-level graphics API
   - More control
   - Steeper learning curve
   - Install: `go get github.com/veandco/go-sdl2/sdl`

## Project Structure

```
KaijuEngine/
├── main.go              # Entry point
├── go.mod               # Go module definition
├── README.md            # Project documentation
├── DEVGUIDE.md          # This file
├── engine/              # Core engine package
│   └── engine.go        # Main game loop
├── game/                # Game logic package
│   └── state.go         # Game state & entities
├── rendering/           # Graphics system package
│   └── renderer.go      # Rendering interface
└── input/               # Input handling package
    └── manager.go       # Input management
```

## Contributing

1. Create a feature branch: `git checkout -b feature/my-feature`
2. Make changes and test
3. Commit with descriptive messages
4. Push and create a pull request

## Code Style

- Follow standard Go conventions
- Use meaningful variable and function names
- Keep functions small and focused
- Add comments for non-obvious logic
- Run `gofmt` before committing: `go fmt ./...`

## Common Tasks

### Add a New Package
```bash
mkdir -p packagename
touch packagename/file.go
```

### Run Code Formatting
```bash
go fmt ./...
goimports -w .  # Requires: go install golang.org/x/tools/cmd/goimports@latest
```

### Check for Issues
```bash
go vet ./...
golint ./...  # Requires: go install github.com/golangci/golint@latest
```

## Next Steps

1. Choose a graphics library (recommend Ebiten)
2. Implement basic sprite rendering
3. Add keyboard input handling
4. Create a simple test game (e.g., Pong or Breakout)
5. Add physics and collision detection

## References

- [Go Documentation](https://golang.org/doc/)
- [Ebiten Documentation](https://ebiten.org/)
- [Game Programming Patterns](https://gameprogrammingpatterns.com/)
