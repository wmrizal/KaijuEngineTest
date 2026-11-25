# KaijuEngine

A basic 2D game engine written in Go with Ebiten graphics support.

## Quick Start

### Build
```bash
go build -o KaijuEngine.exe
```

### Run
```bash
./KaijuEngine.exe
```

A window will appear showing:
- **Animated bouncing circles** in various colors
- **FPS counter** in the top-left
- **Colored shapes** rendering in real-time
- **Blue background** with interactive demo

## Features

✅ **Graphics Rendering** - Ebiten 2D graphics with shapes and colors  
✅ **Game Loop** - 60 FPS target with delta time calculation  
✅ **Entity System** - Basic entity management  
✅ **Input Handling** - Window close and input management  
✅ **Cross-Platform** - Builds on Windows, Linux, macOS  

## Project Structure

```
KaijuEngine/
├── main.go              # Entry point
├── go.mod/go.sum        # Dependencies (Ebiten, etc)
├── engine/              # Core game loop
│   └── engine.go
├── game/                # Game entities & state
│   └── state.go
├── rendering/           # Graphics with Ebiten
│   └── renderer.go
└── input/               # Input management
    └── manager.go
```

## Dependencies

- **Go 1.25+**
- **Ebiten v2** - 2D graphics and game framework
- **Vector drawing** - For shapes and primitives

## Next Steps

1. Add keyboard input handling
2. Create game entities (players, enemies)
3. Implement collision detection
4. Add sprite rendering
5. Create a simple game (Pong, Breakout)

## Documentation

- [QUICKSTART.md](QUICKSTART.md) - 5-minute getting started guide
- [DEVGUIDE.md](DEVGUIDE.md) - Architecture and design patterns

## License

MIT

