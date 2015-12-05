package player

import (
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	actor Actor
}

func NewPlayer(name string, HP, MP, speed, posX, posY int) *Player {
	p := Player{
		actor: *NewActor(name, HP, MP, speed, posX, posY),
	}
	return &p
}

func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.actor.prevX, p.actor.prevY = p.actor.entity.Position()
		prevX := p.actor.prevX
		prevY := p.actor.prevY
		switch event.Key {
		case tl.KeyArrowRight:
			p.actor.entity.SetPosition(prevX+1, prevY)
			break
		case tl.KeyArrowLeft:
			p.actor.entity.SetPosition(prevX-1, prevY)
			break
		case tl.KeyArrowUp:
			p.actor.entity.SetPosition(prevX, prevY-1)
			break
		case tl.KeyArrowDown:
			p.actor.entity.SetPosition(prevX, prevY+1)
			break
		}
	}
}

func (p *Player) Draw(s *tl.Screen) {
	p.actor.Draw(s)
}
