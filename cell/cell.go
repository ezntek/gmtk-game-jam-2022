package cell

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	borderRect   rl.Rectangle
	innerRect    rl.Rectangle
	IsAlive      bool
	IsGenerator  bool
	CellBelogsTo string
}

func (cell Cell) Draw() {
	if !cell.IsAlive {
		rl.DrawRectangleRec(cell.innerRect, rl.LightGray)
	} else {
		switch cell.CellBelogsTo {
		case "player":
			rl.DrawRectangleRec(cell.innerRect, rl.DarkGray)
			if cell.IsGenerator {
				rl.DrawRectangle(int32(cell.innerRect.X+cell.innerRect.Width/4), int32(cell.innerRect.Y+cell.innerRect.Height/4), int32(cell.innerRect.Width/6), int32(cell.innerRect.Height/6), rl.LightGray)
			}
		case "enemy":
			rl.DrawRectangleRec(cell.innerRect, rl.Red)
			if cell.IsGenerator {
				rl.DrawRectangle(int32(cell.innerRect.X+cell.innerRect.Width/4), int32(cell.innerRect.Y+cell.innerRect.Height/4), int32(cell.innerRect.Width/6), int32(cell.innerRect.Height/6), rl.LightGray)
			}
		default:
			break
		}
	}
	rl.DrawRectangleRec(cell.borderRect, color.RGBA{234, 234, 234, 255})
}

func New(atLocation rl.Vector2, size int32) *Cell {
	return &Cell{
		borderRect:   rl.NewRectangle(atLocation.X, atLocation.Y, float32(size-10), float32(size-10)),
		innerRect:    rl.NewRectangle(atLocation.X+5, atLocation.Y+5, float32(size-10), float32(size-10)),
		IsAlive:      false,
		IsGenerator:  false,
		CellBelogsTo: "dead",
	}
}
