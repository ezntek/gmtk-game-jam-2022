package main

import (
	"gmtk_2022/cell"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawGrid(grid *[][]cell.Cell) {
	for _, column := range *grid {
		for _, cell := range column {
			cell.Draw()
		}
	}
}
func CellBelogsTo(grid *[][]cell.Cell, x, y int) string { return (*grid)[y][x+57].CellBelogsTo }
func IsCellAlive(grid *[][]cell.Cell, x, y int) bool    { return (*grid)[y][x+57].IsAlive }
func ChangeCellState(grid *[][]cell.Cell, to bool, x, y int) {
	(*grid)[y][x+57].IsAlive = to
}
func MakeGenerator(grid *[][]cell.Cell, ungeneratorify bool, x, y int) {
	if !ungeneratorify {
		(*grid)[y][x+57].IsGenerator = true
		if !(*grid)[y][x+57].IsAlive {
			(*grid)[y][x+57].IsAlive = true
		}
	} else {
		(*grid)[y][x+57].IsGenerator = false
	}

}
func RollDice(dicAmount int) int {
	if dicAmount >= 1 {
		rand.Seed(time.Now().Local().Unix() + (7 / 3))
		dicAmount--
		return 1 + rand.Intn(6-1)
	}
	return 0
}
func Movement(grid *[][]cell.Cell, generatorCoordinates *rl.Vector2) {
	if rl.IsKeyPressed(rl.KeySpace) {
		ChangeCellState(grid, true, int(generatorCoordinates.X-1), int(generatorCoordinates.Y))
		ChangeCellState(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y))
		ChangeCellState(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y+1))
		ChangeCellState(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y-1))
		ChangeCellState(grid, true, int(generatorCoordinates.X-1), int(generatorCoordinates.Y-1))
		ChangeCellState(grid, true, int(generatorCoordinates.X-1), int(generatorCoordinates.Y+1))
		ChangeCellState(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y-1))
		ChangeCellState(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y+1))
	}
	if rl.IsKeyPressed(rl.KeyLeft) && IsCellAlive(grid, int(generatorCoordinates.X-1), int(generatorCoordinates.Y)) && CellBelogsTo(grid, int(generatorCoordinates.X-1), int(generatorCoordinates.Y)) == "player" {
		MakeGenerator(grid, false, int(generatorCoordinates.X-1), int(generatorCoordinates.Y))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.X -= 1
	}
	if rl.IsKeyPressed(rl.KeyRight) && IsCellAlive(grid, int(generatorCoordinates.X+1), int(generatorCoordinates.Y)) && CellBelogsTo(grid, int(generatorCoordinates.X+1), int(generatorCoordinates.Y)) == "player" {
		MakeGenerator(grid, false, int(generatorCoordinates.X+1), int(generatorCoordinates.Y))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.X += 1
	}
	if rl.IsKeyPressed(rl.KeyUp) && IsCellAlive(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y-1)) && CellBelogsTo(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y-1)) == "player" {
		MakeGenerator(grid, false, int(generatorCoordinates.X), int(generatorCoordinates.Y-1))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.Y -= 1
	}
	if rl.IsKeyPressed(rl.KeyDown) && IsCellAlive(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y+1)) && CellBelogsTo(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y+1)) == "player" {
		MakeGenerator(grid, false, int(generatorCoordinates.X), int(generatorCoordinates.Y+1))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.Y += 1
	}
}

func main() {
	rand.Seed(time.Now().Local().UnixNano())
	generatorCoordinates := rl.Vector2{X: float32(rand.Intn(57)), Y: float32(rand.Intn(38))}
	var screenWidth, screenHeight int32 = 1140, 860
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)
	gridSize := rl.Vector2{X: 57, Y: 38}
	cellWidth := 20

	var mainGrid [][]cell.Cell
	for yptr := 0; yptr < int(gridSize.Y)*cellWidth; yptr += int(cellWidth) {
		mainGrid = append(mainGrid, make([]cell.Cell, int(gridSize.X)))
		for xptr := 0; xptr < int(gridSize.X)*cellWidth; xptr += int(cellWidth) {
			mainGrid[yptr/int(cellWidth)] = append(mainGrid[yptr/int(cellWidth)], cell.New(rl.Vector2{X: float32(xptr), Y: float32(yptr)}, int32(cellWidth)))
		}
	}

	cam := rl.Camera2D{
		Offset:   rl.Vector2{X: 0, Y: 100},
		Rotation: 0,
		Target:   rl.Vector2{X: 0, Y: 0},
		Zoom:     1.0,
	}

	MakeGenerator(&mainGrid, false, int(generatorCoordinates.X), int(generatorCoordinates.Y))

	for !rl.WindowShouldClose() {
		Movement(&mainGrid, &generatorCoordinates)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(cam)
		DrawGrid(&mainGrid)
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
