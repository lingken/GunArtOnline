package main

import (
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()
	// Add a white background
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})

}
