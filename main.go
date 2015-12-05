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
	human := player.NewPlayer("Ken", 100, 100, 100, 10, 5)
	level.AddEntity(human)
	debugInfo := message.NewDebugInfo()
	level.AddEntity(debugInfo)
	game.Screen().SetLevel(level)
	game.Start()
}
