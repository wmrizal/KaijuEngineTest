# KaijuEngine - Quick Start Guide

## 🦖 Welcome to KaijuEngine!

A lightweight 2D game engine in Go with **real graphics rendering** via Ebiten.

## 🎮 What You Get

When you run the demo, you'll see:
- **Animated bouncing circles** in blue/purple
- **FPS counter** showing real-time performance
- **Smooth 60 FPS rendering**
- **Colorful shapes** (squares, circles) in real-time

## ⚡ 2-Minute Setup

### 1. Build
```bash
cd D:\repos\KaijuEngine
go build -o KaijuEngine.exe
```

### 2. Run
```bash
.\KaijuEngine.exe
```

**A window opens with animated graphics!** Press ESC or close the window to exit.

## 📋 What's Inside

| Component | Purpose |
|-----------|---------|
| **engine/** | Game loop & update/render cycle |
| **rendering/** | Graphics with Ebiten |
| **game/** | Entity management |
| **input/** | Input handling |

## 🎨 Rendering Features

✅ Vector drawing (circles, rectangles)  
✅ Color support (RGBA)  
✅ FPS counter  
✅ 60 FPS target  
✅ Smooth animations  

## 💻 Key Code Snippets

### Create the Engine
```go
gameEngine := engine.NewEngine("KaijuEngine", 800, 600)
gameEngine.Run() // Opens window and starts demo
```

### Add an Entity
```go
entity := &game.Entity{
    ID: 1,
    X: 100, Y: 100,
    Width: 32, Height: 32,
    Active: true,
    VelX: 100,  // pixels/sec
    VelY: 0,
}
gameEngine.GetGameState().AddEntity(entity)
```

### Draw Shapes
```go
// In rendering/renderer.go Draw() method:
renderer.DrawFilledCircle(screen, 100, 100, 25, color.RGBA{255, 0, 0, 255})
renderer.DrawFilledRect(screen, 50, 50, 100, 100, color.RGBA{0, 255, 0, 255})
```

## 🚀 Next: Building a Game

### Step 1: Add Input
```go
if inputMgr.IsKeyPressed(ebiten.KeyArrowLeft) {
    player.VelX = -300
}
```

### Step 2: Collision Detection
```go
func CheckCollision(e1, e2 *game.Entity) bool {
    return e1.X < e2.X+e2.Width &&
           e1.X+e1.Width > e2.X &&
           e1.Y < e2.Y+e2.Height &&
           e1.Y+e1.Height > e2.Y
}
```

### Step 3: Create Your Game
Modify `main.go` to add:
- Player entity with keyboard controls
- Enemy entities with AI
- Game logic and win/lose conditions

## 📚 Resources

- **[DEVGUIDE.md](DEVGUIDE.md)** - Architecture deep dive
- **[Ebiten Docs](https://ebiten.org/)** - Graphics library reference
- **[Game Patterns](https://gameprogrammingpatterns.com/)** - Design patterns

## 🎮 Project Ideas

1. **Pong** - 2 paddles, 1 ball (easiest)
2. **Breakout** - Paddle, blocks, ball (medium)
3. **Space Shooter** - Player ship, enemies, bullets (medium)
4. **Platformer** - Player, platforms, gravity (harder)

## 💡 Tips

- Use `rendering/renderer.go` to add new drawing functions
- Use `game/state.go` to add game logic
- Use `input/manager.go` to handle player input
- Run `go fmt ./...` before committing

## ✅ Verification

You know it's working when:
1. ✅ Executable builds without errors
2. ✅ Window appears with blue background
3. ✅ Colored circles animate smoothly
4. ✅ FPS counter displays in top-left
5. ✅ Window closes when you press ESC

---

**Happy game development!** 🎮

For more details, see [DEVGUIDE.md](DEVGUIDE.md)

