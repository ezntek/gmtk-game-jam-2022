package cell

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	borderRect   rl.Rectangle
	InnerRect    rl.Rectangle
	IsAlive      bool
	IsGenerator  bool
	CellBelogsTo string
}

func (cell *Cell) Draw() {
	rl.DrawRectangleRec(cell.borderRect, color.RGBA{234, 234, 234, 255})
	if !cell.IsAlive {
		rl.DrawRectangleRec(cell.InnerRect, rl.LightGray)
	} else {
		switch cell.CellBelogsTo {
		case "player":
			rl.DrawRectangleRec(cell.InnerRect, rl.Lime)
			if cell.IsGenerator {
				rl.DrawRectangle(int32(cell.InnerRect.X+cell.InnerRect.Width/4), int32(cell.InnerRect.Y+cell.InnerRect.Height/4), int32(cell.InnerRect.Width/6), int32(cell.InnerRect.Height/6), rl.LightGray)
			}
		case "enemy":
			rl.DrawRectangleRec(cell.InnerRect, rl.Red)
			if cell.IsGenerator {
				rl.DrawRectangle(int32(cell.InnerRect.X+cell.InnerRect.Width/4), int32(cell.InnerRect.Y+cell.InnerRect.Height/4), int32(cell.InnerRect.Width/6), int32(cell.InnerRect.Height/6), rl.LightGray)
			}
		default:
			break
		}
	}

}

func New(atLocation rl.Vector2, size int32) Cell {
	return Cell{
		borderRect:   rl.NewRectangle(atLocation.X, atLocation.Y, float32(size), float32(size)),
		InnerRect:    rl.NewRectangle(atLocation.X+2, atLocation.Y+2, float32(size-4), float32(size-4)),
		IsAlive:      false,
		IsGenerator:  false,
		CellBelogsTo: "player",
	}
}
