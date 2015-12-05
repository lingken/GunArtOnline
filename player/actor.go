package player

import (
	"GunArtOnline/util"
	tl "github.com/JoelOtter/termloop"
)

type ActorState int

const (
	actorAlive ActorState = iota
	actorDead
)

type Actor struct {
	name      string
	HP        int // health point
	MP        int // mana point
	speed     int
	direction util.Direction
	entity    *tl.Entity
	state     ActorState
	posX      int
	posY      int
}

func NewActor(name string, HP, MP, speed, posX, posY int) *Actor {
	actor := Actor{
		name:   name,
		HP:     HP,
		MP:     MP,
		speed:  speed,
		posX:   posX,
		posY:   posY,
		state:  actorAlive,
		entity: tl.NewEntity(posX, posY, 1, 1),
	}
	return &actor
}
