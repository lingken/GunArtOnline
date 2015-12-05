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
	prevX     int
	prevY     int
}

func NewActor(name string, HP, MP, speed, posX, posY int) *Actor {
	actor := Actor{
		name:   name,
		HP:     HP,
		MP:     MP,
		speed:  speed,
		prevX:  posX,
		prevY:  posY,
		state:  actorAlive,
		entity: tl.NewEntity(posX, posY, 1, 1),
	}
	// use symbol to test
	actor.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Ch: 'M'})
	return &actor
}

func (actor *Actor) Draw(screen *tl.Screen) {
	actor.entity.Draw(screen)
}

func (actor *Actor) Tick(event tl.Event) {

}
