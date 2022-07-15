package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	var screenWidth, screenHeight int32 = 1000, 1000
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}
}
