package main

import (
	"GunArtOnline/message"
	"GunArtOnline/player"
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()
	// Add a white background
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})
	debugInfo := message.NewDebugInfo()
	human := player.NewPlayer("Ken", 100, 100, 100, 10, 5, game, debugInfo)
	level.AddEntity(human)

	level.AddEntity(debugInfo)

	enemy := player.NewEnemy("Enemy", 100, 100, 20, 12, 5, 0)
	level.AddEntity(enemy)
	game.Screen().SetLevel(level)
	game.Start()
}
