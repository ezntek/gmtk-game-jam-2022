package main

import (
	"gmtk_2022/cell"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var enemyList []cell.EnemyGeneratorCell

func DrawGrid(grid *[][]cell.Cell) {
	for _, column := range *grid {
		for _, cell := range column {
			cell.Draw()
		}
	}
}
func CellBelogsTo(grid *[][]cell.Cell, x, y int) string      { return (*grid)[y][x+57].CellBelogsTo }
func IsCellAlive(grid *[][]cell.Cell, x, y int) bool         { return (*grid)[y][x+57].IsAlive }
func ChangeCellState(grid *[][]cell.Cell, to bool, x, y int) { (*grid)[y][x+57].IsAlive = to }
func ChangeCellOwnership(grid *[][]cell.Cell, to string, x, y int) {
	(*grid)[y][x+57].CellBelogsTo = to
}
func MakeGenerator(grid *[][]cell.Cell, ungeneratorify bool, x, y int) {
	if !ungeneratorify {
		(*grid)[y][x+57].IsGenerator = true
		if !(*grid)[y][x+57].IsAlive {
			(*grid)[y][x+57].IsAlive = true
		}
	} else {
		(*grid)[y][x+57].IsGenerator = false
	}

}

func RollDice(diceAmount int) int {
	if diceAmount >= 1 {
		rand.Seed(time.Now().Local().Unix() + (7 / 3))
		diceAmount--
		return 1 + rand.Intn(6-1)
	}
	return 0
}
func Movement(grid *[][]cell.Cell, generatorCoordinates *rl.Vector2, ctr *int) {
	if *ctr%18 == 0 {
		rand.Seed(time.Now().UnixNano())
		enemyList = append(enemyList, cell.NewEnemy())
		//var randval int
		//rv := rand.Intn(4)
		for _, enemy := range enemyList {
			if !(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].EnemyHasSetLocation {
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].EnemyHasSetLocation = true
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = true
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsAlive = true
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].CellBelogsTo = "enemy"
			}
			//(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = true
			//(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsAlive = true
			//(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].CellBelogsTo = "enemy"
			enemy.Update(grid, *generatorCoordinates)
		}
		// left row
		if generatorCoordinates.X-1 >= 0 {
			if generatorCoordinates.Y-1 >= 0 {
				MakeGenerator(grid, true, int(generatorCoordinates.X-1), int(generatorCoordinates.Y-1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X-1), int(generatorCoordinates.Y-1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y)-1)
			}

			MakeGenerator(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X-1), int(generatorCoordinates.Y))
			ChangeCellState(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y))

			if generatorCoordinates.Y+1 <= 37 {
				MakeGenerator(grid, true, int(generatorCoordinates.X-1), int(generatorCoordinates.Y+1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X-1), int(generatorCoordinates.Y+1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y)+1)
			}
		}
		// middle two
		if generatorCoordinates.Y+1 <= 37 {
			MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y+1))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X), int(generatorCoordinates.Y+1))
			ChangeCellState(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y)+1)
		}
		if generatorCoordinates.Y-1 >= 0 {
			MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y-1))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X), int(generatorCoordinates.Y-1))
			ChangeCellState(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y)-1)
		}

		// right row
		if generatorCoordinates.X+1 <= 56 {
			if generatorCoordinates.Y-1 >= 0 {
				MakeGenerator(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y+1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X+1), int(generatorCoordinates.Y+1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y)+1)
			}
			MakeGenerator(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X+1), int(generatorCoordinates.Y))
			ChangeCellState(grid, true, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y))
			if generatorCoordinates.Y+1 <= 37 {
				MakeGenerator(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y-1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X+1), int(generatorCoordinates.Y-1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y)-1)
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) && IsCellAlive(grid, int(generatorCoordinates.X-1), int(generatorCoordinates.Y)) && CellBelogsTo(grid, int(generatorCoordinates.X-1), int(generatorCoordinates.Y)) == "player" && generatorCoordinates.X-1 > 0 {
		MakeGenerator(grid, false, int(generatorCoordinates.X-1), int(generatorCoordinates.Y))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.X -= 1
	}
	if rl.IsKeyPressed(rl.KeyRight) && IsCellAlive(grid, int(generatorCoordinates.X+1), int(generatorCoordinates.Y)) && CellBelogsTo(grid, int(generatorCoordinates.X+1), int(generatorCoordinates.Y)) == "player" && generatorCoordinates.X+1 < 56 {
		MakeGenerator(grid, false, int(generatorCoordinates.X+1), int(generatorCoordinates.Y))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.X += 1
	}
	if rl.IsKeyPressed(rl.KeyUp) && IsCellAlive(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y-1)) && CellBelogsTo(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y-1)) == "player" && generatorCoordinates.Y-1 > 0 {
		MakeGenerator(grid, false, int(generatorCoordinates.X), int(generatorCoordinates.Y-1))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.Y -= 1
	}
	if rl.IsKeyPressed(rl.KeyDown) && IsCellAlive(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y+1)) && CellBelogsTo(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y+1)) == "player" && generatorCoordinates.Y+1 < 37 {
		MakeGenerator(grid, false, int(generatorCoordinates.X), int(generatorCoordinates.Y+1))
		MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y))
		generatorCoordinates.Y += 1
	}
}

type SelectionButton struct {
	border    rl.Rectangle
	Text      string
	Selected  bool
	innerRect rl.Rectangle
}

type Config struct {
	EnemyCount    int
	UpdateSpeed   int
	DiceAmplifier int
}

func main() {
	var screen string = "menu"

	var EnemyCount int = 3
	rand.Seed(time.Now().Local().UnixNano())
	generatorCoordinates := rl.Vector2{X: float32(rand.Intn(57)), Y: float32(rand.Intn(38))}
	var screenWidth, screenHeight int32 = 1140, 860
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)
	gridSize := rl.Vector2{X: 57, Y: 38}
	cellWidth := 20

	var enemies = make([]cell.EnemyGeneratorCell, EnemyCount)
	for i := range enemies {
		enemies[i] = cell.NewEnemy()
	}

	var counter int
	var mainGrid [][]cell.Cell
	for yptr := 0; yptr < int(gridSize.Y)*cellWidth; yptr += int(cellWidth) {
		mainGrid = append(mainGrid, make([]cell.Cell, int(gridSize.X)))
		for xptr := 0; xptr < int(gridSize.X)*cellWidth; xptr += int(cellWidth) {
			mainGrid[yptr/int(cellWidth)] = append(mainGrid[yptr/int(cellWidth)], cell.New(rl.Vector2{X: float32(xptr), Y: float32(yptr)}, int32(cellWidth), "player"))
		}
	}
	/*
		mainGrid[int(enemies[0].AtLocation.Y)][int(enemies[0].AtLocation.Y)+57].IsGenerator = true
		mainGrid[int(enemies[0].AtLocation.Y)][int(enemies[0].AtLocation.Y)+57].CellBelogsTo = "enemy"
		mainGrid[int(enemies[0].AtLocation.Y)][int(enemies[0].AtLocation.Y)+57].IsAlive = true
	*/
	cam := rl.Camera2D{
		Offset:   rl.Vector2{X: 0, Y: 100},
		Rotation: 0,
		Target:   rl.Vector2{X: 0, Y: 0},
		Zoom:     1.0,
	}

	var currentlyConfiguring string

	var motdEntries = [...]string{
		"mechkeys",
		"010909",
		"050209",
		"never gonna give you up, (raylib!)",
		"ze coconut nut is a giant nut",
		"if you eat too much you get very fat!",
		"The Work of easontek2398(tek967)!",
		"Also the work of meowscripty!",
		"Check out Rudelies on spotify!",
		"meow",
		"Do you play mincecraft?",
		"Sad music for life - kittycat",
		"The creators are Gophers...",
		"go.dev",
		"lofi girl will give you hugs",
		"egg is a baldhead",
		"jason is an asshole - egg",
		"eason is an asshole - kittycat",
		"peterguo2009 makes music on soundcloud",
	}

	var configuredDiceAmplifier bool = false
	var doneConfiguringDiceAmplifier bool = false
	var motd = motdEntries[rand.Intn(len(motdEntries))]
	var pressEnterPlayText string = "[Press Enter to play]"
	var diceAmplifierText string = "Dice Amplifier (amplifier for any give dice roll result, press 3 if unsure.) [press d to configure]"
	MakeGenerator(&mainGrid, false, int(generatorCoordinates.X), int(generatorCoordinates.Y))
	var gameConf Config

	for !rl.WindowShouldClose() {
		counter++
		if screen == "menu" {
			if currentlyConfiguring == "" {
				if rl.IsKeyPressed(rl.KeyD) && !configuredDiceAmplifier {
					currentlyConfiguring = "diceamplifier"
				}
			}

			switch currentlyConfiguring {
			case "diceamplifier":
				diceAmplifierText = "Dice Amplifier [press (1) (3) or (5)]"
				if rl.IsKeyPressed(rl.KeyOne) {
					gameConf.DiceAmplifier = 1
					doneConfiguringDiceAmplifier = true
				}
				if rl.IsKeyPressed(rl.KeyThree) {
					doneConfiguringDiceAmplifier = true
					gameConf.DiceAmplifier = 3
				}
				if rl.IsKeyPressed(rl.KeyFive) {
					doneConfiguringDiceAmplifier = true
					gameConf.DiceAmplifier = 5
				}
				if doneConfiguringDiceAmplifier {
					configuredDiceAmplifier = true
					diceAmplifierText = "Dice Amplifier [done]"
					currentlyConfiguring = ""
				}
			}

			if rl.IsKeyPressed(rl.KeyEnter) {
				if gameConf.EnemyCount == 0 && gameConf.DiceAmplifier == 0 {
					pressEnterPlayText = "(Need a config to start! Configure below.)[Press Enter to play]"
				} else {
					screen = "game"
				}
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Conquer Of Tiles", screenWidth/2-180, screenHeight/2-400, 40, rl.DarkGray)
			rl.DrawText(pressEnterPlayText, screenWidth/2-180, screenHeight/2-355, 20, rl.Gray)
			rl.DrawText(motd, screenWidth/2-180, screenHeight/2-310, 20, rl.Gold)
			// options menu
			rl.DrawText("Config", 20, screenHeight/2-20, 50, rl.Black)
			//rl.DrawText(updateSpeedText, screenWidth/2-180, screenHeight/2+65, 20, rl.DarkGray)
			rl.DrawText(diceAmplifierText, 20, screenHeight/2+65, 20, rl.DarkGray)
			rl.EndDrawing()
		}
		if screen == "game" {
			Movement(&mainGrid, &generatorCoordinates, &counter)
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.BeginMode2D(cam)
			DrawGrid(&mainGrid)
			rl.EndMode2D()
			// space for art
			// ---
			rl.EndDrawing()
		}
	}
}
