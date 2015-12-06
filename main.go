package main

import (
	"GunArtOnline/message"
	"GunArtOnline/object"
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()
	// Add a white background
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})
	debugInfo := message.NewDebugInfo()
	human := object.NewPlayer("Ken", 100, 100, 100, 10, 5, game, debugInfo)
	level.AddEntity(human)

	level.AddEntity(debugInfo)

	enemy := object.NewEnemy("Enemy", 5, 100, 5, 12, 5, 0, game, debugInfo)
	level.AddEntity(enemy)
	enemy2 := object.NewEnemy("Enemy", 10, 100, 5, 12, 5, 0, game, debugInfo)
	level.AddEntity(enemy2)
	game.Screen().SetLevel(level)
	game.Start()
}
