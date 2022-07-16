package main

import (
	"fmt"
	"gmtk_2022/grid"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var screenWidth, screenHeight int32 = 1000, 1000
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)

	var cellWidth int32 = 50
	var gridSize int32 = 50
	targetedCell := rl.Vector2{X: 0, Y: 0}
	mainGrid := grid.New(0, 0, cellWidth, int(gridSize))

	cam := rl.Camera2D{
		Offset:   rl.Vector2{X: 0, Y: 0},
		Rotation: 0,
		Target:   rl.Vector2{X: 0, Y: 0},
		Zoom:     1.0,
	}
	//mainGrid.Grid[15][15].IsAlive = true
	for !rl.WindowShouldClose() {
		mainGrid.Grid[3][3].IsAlive = true
		rl.BeginDrawing()
		cam.Target.X = targetedCell.X * float32(cellWidth)
		cam.Target.Y = targetedCell.Y * float32(cellWidth)
		// camera panning
		if rl.IsKeyPressed(rl.KeyLeft) { //left
			if targetedCell.X >= 1 {
				targetedCell.X--
			}
		}
		if rl.IsKeyPressed(rl.KeyRight) { //right
			if targetedCell.X <= float32(gridSize)-1 {
				targetedCell.X++
			}
		}
		if rl.IsKeyPressed(rl.KeyDown) && targetedCell.Y <= float32(gridSize)-1 {
			targetedCell.Y++
		}
		if rl.IsKeyPressed(rl.KeyUp) && targetedCell.Y >= 1 {
			targetedCell.Y--
		}
		fmt.Printf("x %d y %d\n", int(targetedCell.X), int(targetedCell.Y))
		rl.ClearBackground(rl.RayWhite)
		// 2d camera
		rl.BeginMode2D(cam)
		mainGrid.Draw()
		rl.EndMode2D()
		// end 2d camera
		rl.EndDrawing()
	}
}
