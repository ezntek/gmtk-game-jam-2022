package grid

import (
	"gmtk_2022/cell"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	Grid [][]cell.Cell
}

func (g *Grid) Draw() {
	// loop through grid
	for _, column := range g.Grid {
		for _, cell := range column {
			// call draw function within the cell
			cell.Draw()
		}
	}
}

func (g *Grid) ChangeState(to bool, x, y int) {
	g.Grid[y][x].IsAlive = to
}

func New(posx int32, posy int32, cellWidth int32, gridSize rl.Vector2) Grid {
	/*
		Fill an array with cells. xptr and yptr point to the current x and y position it should put a cell at.
	*/

	var gridTmp [][]cell.Cell

	for yptr := 0; yptr < int(int32(gridSize.Y)*cellWidth); yptr += int(cellWidth) {
		gridTmp = append(gridTmp, make([]cell.Cell, int(gridSize.X)))
		for xptr := 0; xptr < int(int32(gridSize.X)*cellWidth); xptr += int(cellWidth) {
			gridTmp[yptr/int(cellWidth)] = append(gridTmp[yptr/int(cellWidth)], cell.New(rl.Vector2{X: float32(xptr), Y: float32(yptr)}, cellWidth))
		}
	}
	return Grid{
		Grid: gridTmp,
	}
}
