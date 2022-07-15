package grid

import (
	"gmtk_2022/cell"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	Grid [][]*cell.Cell
}

func (g Grid) Draw() {
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
		initialise array with `make()`
	*/

	gridTmp := make([][]*cell.Cell, gridSize)
	for i := range gridTmp {
		gridTmp = append(gridTmp, make([]*cell.Cell, gridSize))
	}

	/*
		Fill the array with cells. xptr and yptr point to the current x and y position it should put a cell at.
	*/

	for yptr := 0; yptr < gridSize*int(cellWidth); yptr += int(cellWidth) {
		for xptr := 0; xptr < gridSize*int(cellWidth); xptr += int(cellWidth) {
			gridTmp[xptr][yptr] = cell.New(rl.Vector2{X: float32(xptr), Y: float32(yptr)}, 50)
		}
	}
	return &Grid{
		Grid: gridTmp,
	}
}
