package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "fps-game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update game state here.

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
}
