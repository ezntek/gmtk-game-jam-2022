package main

import (
	"gmtk_2022/grid"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var screenWidth, screenHeight int32 = 1140, 860
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)

	mainGrid := grid.New(0, 0, 20, rl.Vector2{X: 57, Y: 38})

	cam := rl.Camera2D{
		Offset:   rl.Vector2{X: 0, Y: 100},
		Rotation: 0,
		Target:   rl.Vector2{X: 0, Y: 0},
		Zoom:     1.0,
	}

	x, y := rand.Intn(57), rand.Intn(38)
	mainGrid.ChangeState(true, x, y)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		// 2d camera
		rl.BeginMode2D(cam)
		mainGrid.Draw()
		rl.EndMode2D()
		// end 2d camera
		rl.EndDrawing()
	}
}
