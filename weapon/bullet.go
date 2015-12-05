package weapon

import (
	"GunArtOnline/message"
	"GunArtOnline/util"
	tl "github.com/JoelOtter/termloop"
)

type Bullet struct {
	prevX     int
	prevY     int
	symbol    rune
	damage    int
	speed     int
	direction util.Direction
	entity    *tl.Entity
	rangeLeft int
	frame     int
	debug     *message.DebugInfo
	game      *tl.Game
}

func NewBullet(posX, posY, damage, speed, rangeLeft int, direction util.Direction, debug *message.DebugInfo, game *tl.Game) *Bullet {
	bullet := Bullet{
		prevX:     posX,
		prevY:     posY,
		damage:    damage,
		speed:     speed,
		direction: direction,
		rangeLeft: 30,
		entity:    tl.NewEntity(posX, posY, 1, 1),
		debug:     debug,
		frame:     0,
		game:      game,
	}
	bullet.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Ch: 'o'})
	return &bullet
}

func (bullet *Bullet) Draw(s *tl.Screen) {
	if bullet.frame < 2000 {
		bullet.frame += 1
		bullet.entity.Draw(s)
		return
	}
	prevX, prevY := bullet.prevX, bullet.prevY
	switch bullet.direction {
	case util.Up:
		bullet.entity.SetPosition(prevX, prevY-1)
		break
	case util.Down:
		bullet.entity.SetPosition(prevX, prevY+1)
		break
	case util.Left:
		bullet.entity.SetPosition(prevX-1, prevY)
		break
	case util.Right:
		bullet.entity.SetPosition(prevX+1, prevY)
		break
	}
	bullet.prevX, bullet.prevY = bullet.entity.Position()
	bullet.entity.Draw(s)
	bullet.frame = 0
	bullet.rangeLeft -= 1
	if bullet.rangeLeft == 0 {
		s.RemoveEntity(bullet)
	}
}

func (bullet *Bullet) Tick(event tl.Event) {}

func (bullet *Bullet) Position() (int, int) {
	return bullet.entity.Position()
}

func (bullet *Bullet) Size() (int, int) {
	return bullet.entity.Size()
}

func (bullet *Bullet) Collide(collision tl.Physical) {
	bullet.game.Screen().Level().RemoveEntity(bullet)
}
