package game

// Entity represents a game object
type Entity struct {
	ID       int
	X, Y     float64
	Width    float64
	Height   float64
	Active   bool
	VelX     float64
	VelY     float64
}

// GameState holds all game state
type GameState struct {
	entities []*Entity
	frameCount int
}

// NewGameState creates a new game state
func NewGameState() *GameState {
	return &GameState{
		entities:   make([]*Entity, 0),
		frameCount: 0,
	}
}

// AddEntity adds an entity to the game state
func (gs *GameState) AddEntity(entity *Entity) {
	gs.entities = append(gs.entities, entity)
}

// RemoveEntity removes an entity from the game state
func (gs *GameState) RemoveEntity(id int) {
	for i, e := range gs.entities {
		if e.ID == id {
			gs.entities = append(gs.entities[:i], gs.entities[i+1:]...)
			break
		}
	}
}

// Update updates game state
func (gs *GameState) Update(deltaTime float64) {
	gs.frameCount++
	
	// Update all entities
	for _, entity := range gs.entities {
		if entity.Active {
			entity.X += entity.VelX * deltaTime
			entity.Y += entity.VelY * deltaTime
		}
	}
}

// GetEntities returns all entities
func (gs *GameState) GetEntities() []*Entity {
	return gs.entities
}
