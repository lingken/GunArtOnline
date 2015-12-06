package main

import (
	"GunArtOnline/message"
	"GunArtOnline/object"
	"GunArtOnline/util"
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()

	// NumEnemyMutex.Lock()
	// NumPlayerMutex.Lock()
	// build(NumEnemy, NumPlayer, game)
	// NumEnemyMutex.Unlock()
	// NumPlayerMutex.Unlock()
	// Add a white background

	// immitate PAXOS storage
	db := util.NewDatabase()
	reg := util.NewRegisterList()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})
	debugInfo := message.NewDebugInfo(db, reg)
	human := object.NewPlayer("Ken", 12, 100, 100, 25, 10, game, debugInfo)

	level.AddEntity(human)
	db.Put(human.Key, *human)
	reg.Register("Ken")

	level.AddEntity(debugInfo)
	db.Put(debugInfo.Key, *debugInfo)
	reg.Register("debugInfo")

	enemy := object.NewEnemy("Enemy1", 5, 100, 5, 12, 5, 0, game, debugInfo)
	level.AddEntity(enemy)
	db.Put(enemy.Key, *enemy)
	reg.Register("Enemy1")

	enemy2 := object.NewEnemy("Enemy2", 10, 100, 5, 15, 5, 0, game, debugInfo)
	level.AddEntity(enemy2)
	db.Put(enemy2.Key, *enemy2)
	reg.Register("Enemy2")

	game.Screen().SetLevel(level)
	game.Start()
}

// func build(numEnemy, numPlayer int, game *tl.Game) {
// 	numEnemyText := tl.NewText(0, 0, "Number of demons: "+strconv.Itoa(numEnemy),
// 		tl.ColorMagenta, tl.ColorWhite)
// 	game.Screen().AddEntity(numEnemyText)

// 	numPlayerText := tl.NewText(0, 1, "Number of angels: "+strconv.Itoa(numPlayer),
// 		tl.ColorMagenta, tl.ColorWhite)
// 	game.Screen().AddEntity(numPlayerText)
// }
