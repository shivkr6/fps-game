package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Target struct {
	Position rl.Vector3
	Size     rl.Vector3
	Health   int
}

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

	targets := []Target{
		{Position: rl.NewVector3(0, 1, -8), Size: rl.NewVector3(1, 2, 1), Health: 100},
		{Position: rl.NewVector3(-6, 1, -7), Size: rl.NewVector3(1, 2, 1), Health: 100},
		{Position: rl.NewVector3(6, 1, -7), Size: rl.NewVector3(1, 2, 1), Health: 100},
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.DrawPlane(rl.NewVector3(0, 0, 0), rl.NewVector2(20, 20), rl.LightGray)

		// Walls.
		rl.DrawCube(rl.NewVector3(0, 2, -10), 20, 4, 0.5, rl.Gray)
		rl.DrawCube(rl.NewVector3(0, 2, 10), 20, 4, 0.5, rl.Gray)
		rl.DrawCube(rl.NewVector3(-10, 2, 0), 0.5, 4, 20, rl.DarkGray)
		rl.DrawCube(rl.NewVector3(10, 2, 0), 0.5, 4, 20, rl.DarkGray)

		// Obstacles and targets.
		rl.DrawCube(rl.NewVector3(-3, 1, -4), 2, 2, 2, rl.Brown)
		rl.DrawCube(rl.NewVector3(3, 0.5, -6), 2, 1, 2, rl.Brown)
		for _, target := range targets {
			if target.Health > 0 {
				rl.DrawCubeV(target.Position, target.Size, rl.Red)
			}
		}
		rl.EndMode3D()

		cx, cy := int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2)
		rl.DrawLine(cx-6, cy, cx+6, cy, rl.Black)
		rl.DrawLine(cx, cy-6, cx, cy+6, rl.Black)
		rl.DrawText("WASD: move | Mouse: look | Esc: quit", 10, 10, 20, rl.Black)
		rl.EndDrawing()
	}
}
