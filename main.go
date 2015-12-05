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

	enemy := object.NewEnemy("Enemy", 5, 100, 0, 12, 5, 0, game, debugInfo)
	level.AddEntity(enemy)
	game.Screen().SetLevel(level)
	game.Start()
}
