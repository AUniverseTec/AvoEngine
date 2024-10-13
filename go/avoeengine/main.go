package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"go/avoeengine/animation_system"
	"go/avoeengine/input_system"
	"go/avoeengine/lua"
	"go/avoeengine/model"
	"go/avoeengine/physics"
)

func main() {
	// Initialize GLFW
	glfw.Init()

	// Create a window
	window, err := glfw.CreateWindow(800, 600, "Avo Engine", nil, nil)
	if err != nil {
		panic(err)
	}

	// Make the window's context current
	window.MakeContextCurrent()

	// Initialize the animation system
	animation_system.Init()

	// Initialize the input system
	input_system.Init()

	// Initialize the Lua scripting system
	lua.Init()

	// Load the game model
	gameModel := model.Load("models/game_model.obj")

	// Create a physics body for the game model
	body := physics.CreateBody(gameModel)

	// Run the game loop
	for !window.ShouldClose() {
		// Update the animation system
		animation_system.Update()

		// Update the input system
		input_system.Update()

		// Update the Lua scripting system
		lua.Update()

		// Update the physics simulation
		physics.Update()

		// Render the game
		render(window, gameModel, body)

		// Swap buffers and poll for events
		window.SwapBuffers()
		glfw.PollEvents()
	}

	// Clean up
	glfw.Terminate()
}
