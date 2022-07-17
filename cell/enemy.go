package cell

import (
	//"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyGeneratorCell struct {
	determinedPosition   bool
	AtLocation           rl.Vector2
	Direction            int
	directionBufferCount int
}

//func cellBelogsTo(grid *[][]Cell, x, y int) string      { return (*grid)[y][x+57].CellBelogsTo }
//func isCellAlive(grid *[][]Cell, x, y int) bool         { return (*grid)[y][x+57].IsAlive }
func changeCellState(grid *[][]Cell, to bool, x, y int) { (*grid)[y][x+57].IsAlive = to }
func changeCellOwnership(grid *[][]Cell, to string, x, y int) {
	(*grid)[y][x+57].CellBelogsTo = to
}
func makeGenerator(grid *[][]Cell, ungeneratorify bool, x, y int) {
	if !ungeneratorify {
		(*grid)[y][x+57].IsGenerator = true
		if !(*grid)[y][x+57].IsAlive {
			(*grid)[y][x+57].IsAlive = true
		}
	} else {
		(*grid)[y][x+57].IsGenerator = false
	}

}
func (enemy *EnemyGeneratorCell) IsEnemyGeneratorAlive(grid *[][]Cell) bool {
	if (*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].CellBelogsTo == "enemy" {
		return true
	} else {
		return false
	}
}

func (enemy *EnemyGeneratorCell) Update(grid *[][]Cell, playerLocation rl.Vector2) {
	if enemy.IsEnemyGeneratorAlive(grid) {
		//fmt.Println(enemy.directionBufferCount)
		// check player xy

		if (*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57-1].CellBelogsTo == "enemy" && enemy.AtLocation.X-1 > 0 && enemy.AtLocation.X > playerLocation.X {
			// left
			makeGenerator(grid, false, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y))
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y))
			enemy.AtLocation.X -= 1

		} else if (*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57+1].CellBelogsTo == "enemy" && enemy.AtLocation.X+1 < 56 && enemy.AtLocation.X < playerLocation.X {
			makeGenerator(grid, false, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y))
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y))
			enemy.AtLocation.X += 1
		} else if (*grid)[int(enemy.AtLocation.Y)-1][int(enemy.AtLocation.X)+57].CellBelogsTo == "enemy" && enemy.AtLocation.Y-1 > 0 && enemy.AtLocation.Y > playerLocation.Y {
			makeGenerator(grid, false, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1)
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y))
			enemy.AtLocation.Y -= 1
		} else if (*grid)[int(enemy.AtLocation.Y)+1][int(enemy.AtLocation.X)+57].CellBelogsTo == "enemy" && enemy.AtLocation.Y+1 > 37 && enemy.AtLocation.Y < playerLocation.Y {
			makeGenerator(grid, false, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1)
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y))
			enemy.AtLocation.Y += 1
		}

		enemy.MoveGenerator(grid, enemy.Direction)
		//
		if enemy.AtLocation.X-1 >= 0 {
			if enemy.AtLocation.Y-1 >= 0 {
				makeGenerator(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1)
				changeCellState(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1)
			}
			makeGenerator(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y))
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y))
			changeCellState(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y))
			if enemy.AtLocation.Y+1 <= 37 {
				makeGenerator(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1)
				changeCellState(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1)
			}
		}
		// middle 2
		if enemy.AtLocation.Y+1 <= 37 {
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1)
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1)
			changeCellState(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1)
		}
		if enemy.AtLocation.Y-1 >= 0 {
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1)
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1)
			changeCellState(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1)
		}
		// right row
		if enemy.AtLocation.X+1 <= 56 {
			if enemy.AtLocation.Y-1 >= 0 {
				makeGenerator(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1)
				changeCellState(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1)
			}
			makeGenerator(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y))
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y))
			changeCellState(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y))
			if enemy.AtLocation.Y <= 37 {
				makeGenerator(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1)
				changeCellState(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1)
			}
		}
		//enemy.determinedPosition = false
	}
}

func NewEnemy() EnemyGeneratorCell {
	x, y := rand.Intn(57), rand.Intn(38)
	return EnemyGeneratorCell{
		directionBufferCount: 0,
		Direction:            int(rl.GetRandomValue(1, 4)),
		AtLocation:           rl.Vector2{X: float32(x), Y: float32(y)},
	}
}
