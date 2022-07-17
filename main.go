package main

import (
	"fmt"
	"gmtk_2022/cell"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var playerTileCount float32
var enemyTileCount float32
var enemyList []cell.EnemyGeneratorCell
var remainingTurns int = 5

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
		if *ctr%(18*9) == 0 {
			enemyList = append(enemyList, cell.NewEnemy())
		}
		//var randval int
		//rv := rand.Intn(4)
		for i, enemy := range enemyList {
			if !(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].EnemyHasSetLocation {
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].EnemyHasSetLocation = true
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = true
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsAlive = true
				(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].CellBelogsTo = "enemy"
			}
			//(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsGenerator = true
			//(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].IsAlive = true
			//(*grid)[int(enemy.AtLocation.Y)][int(enemy.AtLocation.X)+57].CellBelogsTo = "enemy"

			if !enemy.IsActive {
				enemyList[i] = enemyList[len(enemyList)-1]
				enemyList = enemyList[:len(enemyList)-1]
				break
			}
			enemy.Update(grid, *generatorCoordinates, &enemyTileCount, &playerTileCount)
		}
		// left row
		old := playerTileCount
		if generatorCoordinates.X-1 >= 0 {
			if generatorCoordinates.Y-1 >= 0 {
				if !IsCellAlive(grid, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y)-1) {
					playerTileCount++
				}
				MakeGenerator(grid, true, int(generatorCoordinates.X-1), int(generatorCoordinates.Y-1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X-1), int(generatorCoordinates.Y-1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y)-1)
			}
			if !IsCellAlive(grid, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y)) {
				playerTileCount++
			}
			MakeGenerator(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X-1), int(generatorCoordinates.Y))
			ChangeCellState(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y))

			if generatorCoordinates.Y+1 <= 37 {
				if !IsCellAlive(grid, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y)+1) {
					playerTileCount++
				}
				MakeGenerator(grid, true, int(generatorCoordinates.X-1), int(generatorCoordinates.Y+1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X-1), int(generatorCoordinates.Y+1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)-1, int(generatorCoordinates.Y)+1)
			}
		}
		// middle two
		if generatorCoordinates.Y+1 <= 37 {
			if !IsCellAlive(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y)+1) {
				playerTileCount++
			}
			MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y+1))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X), int(generatorCoordinates.Y+1))
			ChangeCellState(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y)+1)
		}
		if generatorCoordinates.Y-1 >= 0 {
			if !IsCellAlive(grid, int(generatorCoordinates.X), int(generatorCoordinates.Y-1)) {
				playerTileCount++
			}
			MakeGenerator(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y-1))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X), int(generatorCoordinates.Y-1))
			ChangeCellState(grid, true, int(generatorCoordinates.X), int(generatorCoordinates.Y)-1)
		}

		// right row
		if generatorCoordinates.X+1 <= 56 {
			if generatorCoordinates.Y-1 >= 0 {
				if !IsCellAlive(grid, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y)-1) {
					playerTileCount++
				}
				MakeGenerator(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y+1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X+1), int(generatorCoordinates.Y+1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y)+1)
			}
			if !IsCellAlive(grid, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y)) {
				playerTileCount++
			}
			MakeGenerator(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y))
			ChangeCellOwnership(grid, "player", int(generatorCoordinates.X+1), int(generatorCoordinates.Y))
			ChangeCellState(grid, true, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y))
			if generatorCoordinates.Y+1 <= 37 {
				if !IsCellAlive(grid, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y)+1) {
					playerTileCount++
				}
				MakeGenerator(grid, true, int(generatorCoordinates.X+1), int(generatorCoordinates.Y-1))
				ChangeCellOwnership(grid, "player", int(generatorCoordinates.X+1), int(generatorCoordinates.Y-1))
				ChangeCellState(grid, true, int(generatorCoordinates.X)+1, int(generatorCoordinates.Y)-1)
			}
		}
		if playerTileCount > old {
			remainingTurns--
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

type Config struct {
	//EnemyCount    int
	//UpdateSpeed   int
	DiceAmplifier int
}

func main() {
	var screen string = "menu"
	rand.Seed(time.Now().Local().UnixNano())
	generatorCoordinates := rl.Vector2{X: float32(rand.Intn(57)), Y: float32(rand.Intn(38))}
	var screenWidth, screenHeight int32 = 1140, 860
	rl.InitWindow(screenWidth, screenHeight, "GMTK JAM 2022")
	rl.SetTargetFPS(60)
	gridSize := rl.Vector2{X: 57, Y: 38}
	cellWidth := 20

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
	var label string
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
				if gameConf.DiceAmplifier == 0 {
					pressEnterPlayText = "(Need a config to start! Configure below.)[Press Enter to play]"
				} else {
					screen = "game"
				}
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Conquerer Of Tiles", screenWidth/2-180, screenHeight/2-400, 40, rl.DarkGray)
			rl.DrawText(pressEnterPlayText, screenWidth/2-180, screenHeight/2-355, 20, rl.Gray)
			rl.DrawText(motd, screenWidth/2-180, screenHeight/2-310, 20, rl.Gold)
			// options menu
			rl.DrawText("Config", 20, screenHeight/2-20, 50, rl.Black)
			//rl.DrawText(updateSpeedText, screenWidth/2-180, screenHeight/2+65, 20, rl.DarkGray)
			rl.DrawText(diceAmplifierText, 20, screenHeight/2+65, 20, rl.DarkGray)
			rl.EndDrawing()
		}
		if screen == "game" {
			if remainingTurns < 1 {
				screen = "death"
			}
			if playerTileCount*100/(gridSize.X*gridSize.Y) > 50 {
				screen = "governmentdestroyed"
			}
			if rl.IsKeyPressed(rl.KeyEnter) {
				a := rand.Intn(20)
				remainingTurns += a
				label = fmt.Sprintf("You got %d more spreads!", a)
			}

			Movement(&mainGrid, &generatorCoordinates, &counter)
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.BeginMode2D(cam)
			DrawGrid(&mainGrid)
			rl.EndMode2D()
			rl.DrawRectangle(20, 20, 110, 30, rl.Black)
			rl.DrawRectangle(25, 25, 100, 20, rl.RayWhite)
			rl.DrawRectangle(25, 25, (int32(playerTileCount) * 100 / (int32(gridSize.X) * int32(gridSize.Y))), 20, rl.SkyBlue)
			rl.DrawRectangle(20, 60, 110, 30, rl.Black)
			rl.DrawRectangle(25, 65, 100, 20, rl.RayWhite)
			rl.DrawRectangle(25, 65, (int32(enemyTileCount) * 100 / (int32(gridSize.X) * int32(gridSize.Y))), 20, rl.Pink)
			rl.DrawText(fmt.Sprintf("%d spreads left [%s]", remainingTurns, label), 140, 20, 20, rl.Gray)
			rl.EndDrawing()
		}
		if screen == "governmentdestroyed" {
			if rl.IsKeyPressed(rl.KeyEnter) {
				screen = "title"
			}
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("GOVERNMENT DESTROYED.", 30, 30, 60, rl.Black)
			rl.DrawText("You conquered more than 50 percent of the map. congrats! Press [enter] to go back to title screen", 30, 100, 20, rl.Gray)
			rl.EndDrawing()
		}
		if screen == "death" {
			if rl.IsKeyPressed(rl.KeyEnter) {
				screen = "title"
			}
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("You died...", 30, 30, 50, rl.Black)
			rl.DrawText("Press [enter] to play again", 30, 90, 20, rl.Gray)
			rl.EndDrawing()
		}
	}
}
