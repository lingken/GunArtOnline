package object

import (
	// "GunArtOnline"
	"GunArtOnline/message"
	"GunArtOnline/util"
	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	Actor
}

func NewPlayer(name string, HP, MP, speed, posX, posY int, game *tl.Game, debug *message.DebugInfo) *Player {
	p := Player{
		Actor: *NewActor(name, HP, MP, speed, posX, posY, game, debug),
	}
	p.direction = util.Right
	return &p
}

func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.prevX, p.prevY = p.entity.Position()
		prevX := p.prevX
		prevY := p.prevY
		switch event.Key {
		case tl.KeyArrowRight:
			p.entity.SetPosition(prevX+1, prevY)
			p.direction = util.Right
			break
		case tl.KeyArrowLeft:
			p.entity.SetPosition(prevX-1, prevY)
			p.direction = util.Left
			break
		case tl.KeyArrowUp:
			p.entity.SetPosition(prevX, prevY-1)
			p.direction = util.Up
			break
		case tl.KeyArrowDown:
			p.entity.SetPosition(prevX, prevY+1)
			p.direction = util.Down
			break
		case tl.KeySpace:
			// posX, posY, damage, speed, rangeLeft int, direction util.Direction
			x, y := p.entity.Position()
			bX, bY := x, y
			switch p.direction {
			case util.Up:
				bY -= 1
				break
			case util.Down:
				bY += 1
				break
			case util.Left:
				bX -= 1
				break
			case util.Right:
				bX += 1
				break
			}
			bullet := NewBullet(bX, bY, 1, 200, 10, p.direction, p.debug, p.game)
			// p.game.Screen().AddEntity(bullet)
			p.game.Screen().Level().AddEntity(bullet)
		}
	}
}
