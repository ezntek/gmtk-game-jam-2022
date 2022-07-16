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
				rl.DrawRectangle(cell.InnerRect.ToInt32().X+2, cell.InnerRect.ToInt32().Y+2, 7, 7, rl.DarkGreen)
			}
		case "enemy":
			rl.DrawRectangleRec(cell.InnerRect, rl.Pink)
			if cell.IsGenerator {
				rl.DrawRectangle(cell.InnerRect.ToInt32().X+2, cell.InnerRect.ToInt32().Y+2, 7, 7, rl.Red)
			}
		default:
			break
		}
	}

}

func New(atLocation rl.Vector2, size int32, as string) Cell {
	return Cell{
		borderRect:   rl.NewRectangle(atLocation.X, atLocation.Y, float32(size), float32(size)),
		InnerRect:    rl.NewRectangle(atLocation.X+1, atLocation.Y+1, float32(size-2), float32(size-2)),
		IsAlive:      false,
		IsGenerator:  false,
		CellBelogsTo: as,
	}
}
