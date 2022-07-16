package cell

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyGeneratorCell struct {
	AtLocation rl.Vector2
	Direction  int
}

func (enemy EnemyGeneratorCell) Update() {
	rand.Seed(time.Now().Local().UnixNano()/int64(enemy.Direction) + (7 / 6))
	enemy.Direction = rand.Intn(4)

	switch enemy.Direction {
	case 1: // left

	case 2:
		break
	case 3:
		break
	case 4:
		break
	default:
		break
	}
}

func NewEnemy(atLocation rl.Vector2, size int32) EnemyGeneratorCell {
	x, y := 57+rand.Intn(114-57), rand.Intn(38)
	return EnemyGeneratorCell{
		AtLocation: rl.Vector2{X: float32(x), Y: float32(y)},
	}
}
