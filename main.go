package main

import (
	"fmt"
	"gmtk_2022/grid"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Vector2x2 struct {
	X rl.Vector2
	Y rl.Vector2
}

func main() {
	var screenWidth, screenHeight int32 = 600, 800
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)
	targetedCell := rl.Vector2{X: 0, Y: 0}
	var cellWidth int32 = 25
	var gridSize int32 = 50

	var renderedArea Vector2x2 = Vector2x2{
		X: targetedCell,
		Y: rl.Vector2{X: targetedCell.X + float32(screenWidth/cellWidth), Y: targetedCell.Y + float32(screenHeight/cellWidth)},
	}

	mainGrid := grid.New(0, 0, cellWidth, int(gridSize))

	cam := rl.Camera2D{
		Offset:   rl.Vector2{X: 0, Y: 0},
		Rotation: 0,
		Target:   rl.Vector2{X: 0, Y: 0},
		Zoom:     1.0,
	}
	//mainGrid.Grid[15][15].IsAlive = true
	for !rl.WindowShouldClose() {
		cam.Offset.Y = 100
		renderedArea.X = targetedCell
		renderedArea.Y = rl.Vector2{X: targetedCell.X + float32(screenWidth/cellWidth), Y: targetedCell.Y + float32(screenHeight/cellWidth)}
		//mainGrid.Grid[3][3].IsAlive = true
		rl.BeginDrawing()
		cam.Target.X = targetedCell.X * float32(cellWidth)
		cam.Target.Y = targetedCell.Y * float32(cellWidth)
		// camera panning
		if rl.IsKeyPressed(rl.KeyZ) { //left
			if targetedCell.X >= 1 {
				targetedCell.X--
			}
		}
		if rl.IsKeyPressed(rl.KeyC) { //right
			if targetedCell.X <= 25 {
				targetedCell.X++
			}
		}
		if rl.IsKeyPressed(rl.KeyX) && targetedCell.Y <= 21 { //down
			targetedCell.Y++
		}
		if rl.IsKeyPressed(rl.KeyR) && targetedCell.Y >= 1 { //up
			targetedCell.Y--
		}
		//fmt.Printf("x %d y %d\n", int(targetedCell.X), int(targetedCell.Y))
		rl.ClearBackground(rl.RayWhite)
		// 2d camera
		rl.BeginMode2D(cam)
		mainGrid.Draw()
		rl.EndMode2D()
		// end 2d camera
		rl.DrawText(fmt.Sprintf("Onscreen: (%d, %d) to (%d, %d)", int(renderedArea.X.X), int(renderedArea.X.Y), int(renderedArea.Y.X), int(renderedArea.Y.Y)), 20, 20, 30, rl.Gray)

		rl.EndDrawing()
	}
}
