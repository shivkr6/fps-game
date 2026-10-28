package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "fps-game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	camera := rl.Camera3D{
		Position:   rl.NewVector3(0, 1.8, 0),
		Target:     rl.NewVector3(0, 1.8, -1),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       60,
		Projection: rl.CameraPerspective,
	}
	rl.DisableCursor()

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
