package cell

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyGeneratorCell struct {
	AtLocation rl.Vector2
	Direction  int
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
func (enemy *EnemyGeneratorCell) IsEnemyGeneratorAlive(grid *[][]Cell, x, y int) bool {
	if (*grid)[y][x+57].CellBelogsTo == "enemy" && (*grid)[y][x+57].IsGenerator {
		return true
	} else {
		return false
	}
}

func (enemy *EnemyGeneratorCell) MoveGenerator(grid *[][]Cell, direction string) {
	switch direction {
	case "left":
		if (*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+56].CellBelogsTo == "enemy" && enemy.AtLocation.X-1 > 0 {
			(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+56].IsGenerator = true
			(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = false
			enemy.AtLocation.X -= 1
		}

	case "right":
		if (*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+58].CellBelogsTo == "enemy" && enemy.AtLocation.X+1 < 56 {
			(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+58].IsGenerator = true
			(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = false
			enemy.AtLocation.X += 1
		}

	case "up":
		if (*grid)[int(enemy.AtLocation.Y)-1][int(enemy.AtLocation.X)+57].CellBelogsTo == "enemy" && enemy.AtLocation.Y-1 > 0 {
			(*grid)[int(enemy.AtLocation.Y)-1][int(enemy.AtLocation.X)+57].IsGenerator = true
			(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = false
			enemy.AtLocation.Y -= 1
		}

	case "down":
		if (*grid)[int(enemy.AtLocation.Y)+1][int(enemy.AtLocation.X)+57].CellBelogsTo == "enemy" && enemy.AtLocation.Y+1 > 37 {
			(*grid)[int(enemy.AtLocation.Y)+1][int(enemy.AtLocation.X)+57].IsGenerator = true
			(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = false
			enemy.AtLocation.Y += 1
		}

	}
}

func (enemy *EnemyGeneratorCell) Update(grid *[][]Cell, randval int) {
	if enemy.IsEnemyGeneratorAlive(grid, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)) {
		fmt.Println(randval)
		enemy.Direction = randval
		switch enemy.Direction {
		case 1: // left
			enemy.MoveGenerator(grid, "left")
		case 2: // right
			enemy.MoveGenerator(grid, "right")
		case 3: // up
			enemy.MoveGenerator(grid, "up")
		case 4: // down
			enemy.MoveGenerator(grid, "down")
		default: // return
			return
		}
		// left row
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
	}
}

func NewEnemy() EnemyGeneratorCell {
	x, y := rand.Intn(57), rand.Intn(38)
	return EnemyGeneratorCell{
		Direction:  int(rl.GetRandomValue(1, 32)),
		AtLocation: rl.Vector2{X: float32(x), Y: float32(y)},
	}
}
