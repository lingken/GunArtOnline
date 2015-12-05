package player

import (
	tl "github.com/JoelOtter/termloop"
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
