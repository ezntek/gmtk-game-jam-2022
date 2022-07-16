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

func New(posx int32, posy int32, cellWidth int32, gridSize int) *Grid {
	/*
		Fill an array with cells. xptr and yptr point to the current x and y position it should put a cell at.
	*/

	var gridTmp [][]cell.Cell

	for yptr := 0; yptr < gridSize*int(cellWidth); yptr += int(cellWidth) {
		gridTmp = append(gridTmp, make([]cell.Cell, gridSize))
		for xptr := 0; xptr < gridSize*int(cellWidth); xptr += int(cellWidth) {
			gridTmp[yptr/int(cellWidth)] = append(gridTmp[yptr/int(cellWidth)], cell.New(rl.Vector2{X: float32(xptr), Y: float32(yptr)}, 50))
		}
	}
	return &Grid{
		Grid: gridTmp,
	}
}
