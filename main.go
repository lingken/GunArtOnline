package main

import (
	"GunArtOnline/message"
	"GunArtOnline/object"
	"GunArtOnline/util"
	"fmt"
	// "github.com/cmu440-F15/paxosapp"
	// "net/rpc"
	"encoding/gob"
	"os"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	// Check sufficient argument
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <host:port> <username>")
		return
	}

	// Connect the game to PaxosNode with specified hostport
	paxos := util.NewPaxos(os.Args[1])

	gob.Register(object.Actor{})
	gob.Register(object.Bullet{})
	gob.Register(object.Enemy{})
	gob.Register(object.Player{})
	// gob.Register(tl.Text{})
	gob.Register(tl.Entity{})
	gob.Register(tl.Canvas{})
	gob.Register(tl.Cell{})
	// gob.Register(tl.Attr{})

	// fmt.Println(paxos)

	game := tl.NewGame()

	// immitate PAXOS storage
	db := util.NewDatabase(paxos)
	reg := util.NewRegisterList()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})
	debugInfo := message.NewDebugInfo(db, reg)

	level.AddEntity(debugInfo)
	db.Put(debugInfo.Key, *debugInfo)
	reg.Register("debugInfo")

	human := object.NewPlayer(os.Args[2], 12, 100, 100, 25, 10, game, debugInfo, db, reg)

	level.AddEntity(human)
	db.Put(human.Key, *human)
	reg.Register("Ken")

	enemy := object.NewEnemy("Enemy1", 5, 100, 0, 12, 5, 0, game, debugInfo, db, reg)
	level.AddEntity(enemy)
	db.Put(enemy.Key, *enemy)
	reg.Register("Enemy1")

	enemy2 := object.NewEnemy("Enemy2", 10, 100, 0, 15, 5, 0, game, debugInfo, db, reg)
	level.AddEntity(enemy2)
	db.Put(enemy2.Key, *enemy2)
	reg.Register("Enemy2")

	game.Screen().SetLevel(level)
	game.Start()
}
