package game

import (
	"go/avoeengine/model"
	"go/avoeengine/physics"
)

type Game struct {
	Model *model.Model
	Body  *physics.Body
}

func New() *Game {
	// Create a new game instance
	// ...
	return &Game{}
}
