package cell

import (
	//"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyGeneratorCell struct {
	IsActive   bool
	AtLocation rl.Vector2
	Direction  int
}

func isCellAlive(grid *[][]Cell, x, y int) bool    { return (*grid)[y][x+57].IsAlive }
func cellBelogsTo(grid *[][]Cell, x, y int) string { return (*grid)[y][x+57].CellBelogsTo }

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

func (enemy *EnemyGeneratorCell) Update(grid *[][]Cell, playerLocation rl.Vector2, enemyTiles *float32, playerTileCount *float32) {
	enemy.IsActive = enemy.IsEnemyGeneratorAlive(grid)
	if enemy.IsActive {
		// misc movement things

		// left row
		if enemy.AtLocation.X-1 >= 0 {
			if enemy.AtLocation.Y-1 >= 0 {
				if !isCellAlive(grid, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1) {
					*enemyTiles++
					if cellBelogsTo(grid, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1) == "player" {
						*playerTileCount--
					}
				}
				makeGenerator(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1)
				changeCellState(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)-1)
			}
			if !isCellAlive(grid, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)) {
				*enemyTiles++
				if cellBelogsTo(grid, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)) == "player" {
					*playerTileCount--
				}
			}
			makeGenerator(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y))
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y))
			changeCellState(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y))
			if enemy.AtLocation.Y+1 <= 37 {
				if !isCellAlive(grid, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1) {
					*enemyTiles++
					if cellBelogsTo(grid, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1) == "player" {
						*playerTileCount--
					}
				}
				makeGenerator(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1)
				changeCellState(grid, true, int(enemy.AtLocation.X)-1, int(enemy.AtLocation.Y)+1)
			}
		}
		// middle 2
		if enemy.AtLocation.Y+1 <= 37 {
			if !isCellAlive(grid, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1) {
				*enemyTiles++
				if cellBelogsTo(grid, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1) == "player" {
					*playerTileCount--
				}
			}
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1)
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1)
			changeCellState(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)+1)
		}
		if enemy.AtLocation.Y-1 >= 0 {
			if !isCellAlive(grid, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1) {
				*enemyTiles++
				if cellBelogsTo(grid, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1) == "player" {
					*playerTileCount--
				}
			}
			makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1)
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1)
			changeCellState(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y)-1)
		}
		// right row
		if enemy.AtLocation.X+1 <= 56 {
			if enemy.AtLocation.Y-1 >= 0 {
				if !isCellAlive(grid, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1) {
					*enemyTiles++
					if cellBelogsTo(grid, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1) == "player" {
						*playerTileCount--
					}
				}
				makeGenerator(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1)
				changeCellState(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1)
			}
			if !isCellAlive(grid, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)-1) {
				*enemyTiles++
				if cellBelogsTo(grid, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)) == "player" {
					*playerTileCount--
				}
			}
			makeGenerator(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y))
			changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y))
			changeCellState(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y))
			if enemy.AtLocation.Y <= 37 {
				if !isCellAlive(grid, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1) {
					*enemyTiles++
					if cellBelogsTo(grid, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1) == "player" {
						*playerTileCount--
					}
				}
				makeGenerator(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1)
				changeCellOwnership(grid, "enemy", int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1)
				changeCellState(grid, true, int(enemy.AtLocation.X)+1, int(enemy.AtLocation.Y)+1)
			}
		}
		makeGenerator(grid, true, int(enemy.AtLocation.X), int(enemy.AtLocation.Y))
		//enemy.determinedPosition = false
	}
}

func NewEnemy() EnemyGeneratorCell {
	x, y := 1+rand.Intn(56-1), 1+rand.Intn(37-1)
	return EnemyGeneratorCell{
		IsActive:   true,
		Direction:  int(rl.GetRandomValue(1, 4)),
		AtLocation: rl.Vector2{X: float32(x), Y: float32(y)},
	}
}
