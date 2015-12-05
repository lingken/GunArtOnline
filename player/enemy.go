package player

import (
	"GunArtOnline/util"
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

type Enemy struct {
	Actor
	hatred int
	target []*Actor
}

func NewEnemy(name string, HP, MP, speed, posX, posY, hatred int) *Enemy {
	enemy := Enemy{
		Actor:  *NewActor(name, HP, MP, speed, posX, posY),
		hatred: hatred,
		target: make([]*Actor, 0),
	}
	enemy.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Ch: 'E'})
	return &enemy
}

func (enemy *Enemy) Draw(s *tl.Screen) {
	if enemy.speed == 0 || enemy.frame < util.Timeconst/enemy.speed {
		enemy.frame += 1
		enemy.entity.Draw(s)
		return
	}
	enemy.prevX, enemy.prevY = enemy.entity.Position()
	prevX, prevY := enemy.prevX, enemy.prevY
	switch rand.Int() % 4 {
	case 0: // Up
		enemy.entity.SetPosition(prevX, prevY-1)
		enemy.direction = util.Up
		break
	case 1: // Down
		enemy.entity.SetPosition(prevX, prevY+1)
		enemy.direction = util.Down
		break
	case 2: // Left
		enemy.entity.SetPosition(prevX-1, prevY)
		enemy.direction = util.Left
		break
	case 3: // Right
		enemy.entity.SetPosition(prevX+1, prevY)
		enemy.direction = util.Right
		break
	}
	enemy.frame = 0
}
