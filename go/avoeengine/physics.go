package physics

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Body struct {
	Position mgl32.Vec3
	Velocity mgl32.Vec3
}

func CreateBody(model *model.Model) *Body {
	// Create a physics body for the model
	// ...
	return &Body{}
}

func Update() {
	// Update the physics simulation
	// ...
}
