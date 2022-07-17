package cell

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	EnemyHasSetLocation bool
	borderRect          rl.Rectangle
	InnerRect           rl.Rectangle
	IsAlive             bool
	IsGenerator         bool
	CellBelogsTo        string
}

func (cell *Cell) Draw() {
	rl.DrawRectangleRec(cell.borderRect, color.RGBA{234, 234, 234, 255})
	if !cell.IsAlive {
		rl.DrawRectangleRec(cell.InnerRect, rl.LightGray)
	} else {
		switch cell.CellBelogsTo {
		case "player":
			rl.DrawRectangleRec(cell.InnerRect, rl.SkyBlue)
			if cell.IsGenerator {
				rl.DrawRectangle(cell.InnerRect.ToInt32().X+2, cell.InnerRect.ToInt32().Y+2, (cell.InnerRect.ToInt32().Width/2)-2, (cell.InnerRect.ToInt32().Height/2)-2, rl.DarkBlue)
				//rl.DrawTextureRec(cell.PlayerGenTexture, cell.borderRect, rl.Vector2{X: cell.InnerRect.X - 1, Y: cell.InnerRect.Y - 1}, rl.RayWhite)
			}
		case "enemy":
			rl.DrawRectangleRec(cell.InnerRect, rl.Pink)
			if cell.IsGenerator {
				rl.DrawRectangle(cell.InnerRect.ToInt32().X+2, cell.InnerRect.ToInt32().Y+2, (cell.InnerRect.ToInt32().Width/2)-2, (cell.InnerRect.ToInt32().Height/2)-2, rl.Red)
				//rl.DrawTextureRec(cell.EnemyGenTexture, cell.InnerRect, rl.Vector2{X: cell.InnerRect.X, Y: cell.InnerRect.Y + 1}, rl.RayWhite)
			}
		default:
			break
		}
	}

}

func New(atLocation rl.Vector2, size int32, as string) Cell {
	return Cell{
		EnemyHasSetLocation: false,
		borderRect:          rl.NewRectangle(atLocation.X, atLocation.Y, float32(size), float32(size)),
		InnerRect:           rl.NewRectangle(atLocation.X+1, atLocation.Y+1, float32(size-2), float32(size-2)),
		IsAlive:             false,
		IsGenerator:         false,
		CellBelogsTo:        "player",
	}
}
