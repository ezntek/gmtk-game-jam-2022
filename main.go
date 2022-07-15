package main

import (
	"gmtk_2022/grid"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var screenWidth, screenHeight int32 = 1000, 1000
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)

	myGrid := grid.New(0, 0, 50, 5)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		myGrid.Draw()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
}
