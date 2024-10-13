package model

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Model struct {
	Vertices []mgl32.Vec3
	Indices  []uint32
	Material *Material
}

func Load(path string) *Model {
	// Load the model from the file
	// ...
	return &Model{}
}
