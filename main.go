package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Position rl.Vector3
	Size     rl.Vector3
	Color    rl.Color
}

type Target struct {
	Position rl.Vector3
	Size     rl.Vector3
	Health   int
}

func boxBounds(position, size rl.Vector3) rl.BoundingBox {
	half := rl.Vector3Scale(size, 0.5)
	return rl.BoundingBox{
		Min: rl.Vector3Subtract(position, half),
		Max: rl.Vector3Add(position, half),
	}
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

	// Shared geometry for drawing and shot blocking.
	blocks := []Block{
		{rl.NewVector3(0, 2, -10), rl.NewVector3(20, 4, 0.5), rl.Gray},
		{rl.NewVector3(0, 2, 10), rl.NewVector3(20, 4, 0.5), rl.Gray},
		{rl.NewVector3(-10, 2, 0), rl.NewVector3(0.5, 4, 20), rl.DarkGray},
		{rl.NewVector3(10, 2, 0), rl.NewVector3(0.5, 4, 20), rl.DarkGray},
		{rl.NewVector3(-3, 1, -4), rl.NewVector3(2, 2, 2), rl.Brown},
		{rl.NewVector3(3, 0.5, -6), rl.NewVector3(2, 1, 2), rl.Brown},
	}

	targets := []Target{
		{Position: rl.NewVector3(0, 1, -8), Size: rl.NewVector3(1, 2, 1), Health: 100},
		{Position: rl.NewVector3(-6, 1, -7), Size: rl.NewVector3(1, 2, 1), Health: 100},
		{Position: rl.NewVector3(6, 1, -7), Size: rl.NewVector3(1, 2, 1), Health: 100},
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			center := rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)
			ray := rl.GetScreenToWorldRay(center, camera)
			nearest := float32(math.MaxFloat32)
			hitTarget := -1

			for _, block := range blocks {
				hit := rl.GetRayCollisionBox(ray, boxBounds(block.Position, block.Size))
				if hit.Hit && hit.Distance < nearest {
					nearest = hit.Distance
				}
			}
			for i, target := range targets {
				if target.Health <= 0 {
					continue
				}
				hit := rl.GetRayCollisionBox(ray, boxBounds(target.Position, target.Size))
				if hit.Hit && hit.Distance < nearest {
					nearest = hit.Distance
					hitTarget = i
				}
			}
			if hitTarget >= 0 {
				targets[hitTarget].Health -= 25
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.DrawPlane(rl.NewVector3(0, 0, 0), rl.NewVector2(20, 20), rl.LightGray)

		for _, block := range blocks {
			rl.DrawCubeV(block.Position, block.Size, block.Color)
		}
		for _, target := range targets {
			if target.Health > 0 {
				rl.DrawCubeV(target.Position, target.Size, rl.Red)
			}
		}
		rl.EndMode3D()

		cx, cy := int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2)
		rl.DrawLine(cx-6, cy, cx+6, cy, rl.Black)
		rl.DrawLine(cx, cy-6, cx, cy+6, rl.Black)
		rl.DrawText("WASD: move | Mouse: look | LMB: shoot | Esc: quit", 10, 10, 20, rl.Black)
		rl.EndDrawing()
	}
}
