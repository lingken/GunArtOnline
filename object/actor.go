package object

import (
	"GunArtOnline/message"
	"GunArtOnline/util"
	"fmt"
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
	frame     int
	direction util.Direction
	entity    *tl.Entity
	state     ActorState
	prevX     int
	prevY     int
	game      *tl.Game
	debug     *message.DebugInfo
}

func NewActor(name string, HP, MP, speed, posX, posY int, game *tl.Game, debug *message.DebugInfo) *Actor {
	actor := Actor{
		name:   name,
		HP:     HP,
		MP:     MP,
		speed:  speed,
		prevX:  posX,
		prevY:  posY,
		state:  actorAlive,
		entity: tl.NewEntity(posX, posY, 1, 1),
		game:   game,
		debug:  debug,
	}
	// use symbol to test
	actor.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Ch: 'M'})
	return &actor
}

func (actor *Actor) Hit(bullet *Bullet) {
	// actor.HP -= bullet.Damage
	actor.debug.AddInfo(fmt.Sprintf("Abstract actor Hit remain HP: %d\n", actor.HP))
	// if actor.HP <= 0 {
	// 	actor.game.Screen().Level().RemoveEntity(actor)
	// }
}

func (actor *Actor) Draw(screen *tl.Screen) {
	actor.entity.Draw(screen)
}

func (actor *Actor) Tick(event tl.Event) {

}

func (actor *Actor) Position() (int, int) {
	return actor.entity.Position()
}

func (actor *Actor) Size() (int, int) {
	return actor.entity.Size()
}

func (actor *Actor) Collide(collision tl.Physical) {
	// if _, ok := collision.(*tl.Rectangle); ok {
	// 	actor.entity.SetPosition(actor.prevX, actor.prevY)
	// 	// or if it is another mech
	// } else if _, ok := collision.(*Enemy); ok {
	// 	actor.entity.SetPosition(actor.prevX, actor.prevY)
	// }
	// x, y := collision.Position()
	// nx, ny := actor.entity.Position()
	// if x == nx && y == ny {
	// actor.entity.SetPosition(actor.prevX, actor.prevY)
	// }
}
