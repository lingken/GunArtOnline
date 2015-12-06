package object

import (
	// "GunArtOnline"
	"GunArtOnline/message"
	"GunArtOnline/util"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	//"strconv"
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

func (p *Player) Hit(bullet *Bullet) {
	// only bullets from enemy can hurt player
	if bullet.sourceType == demon {
		p.HP -= bullet.Damage
		if v, ok := bullet.source.(*Enemy); ok {
			p.debug.AddInfo(fmt.Sprintf("Player Hit remain HP: %d by %s\n", p.HP, v.name))
		} else {
			p.debug.AddInfo(fmt.Sprintf("Player Hit remain HP: %d\n", p.HP))
		}

		if p.HP <= 0 {
			p.state = actorDead
			p.game.Screen().Level().RemoveEntity(p)

			// NumPlayerMutex.Lock()
			// NumPlayer -= 1
			// numPlayerText := tl.NewText(0, 1, "Number of angels: "+strconv.Itoa(NumPlayer),
			// 	tl.ColorMagenta, tl.ColorWhite)
			// p.game.Screen().AddEntity(numPlayerText)
			// NumPlayerMutex.Unlock()
		}
	}
}

func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.prevX, p.prevY = p.entity.Position()
		prevX, prevY := p.entity.Position()
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
			bullet := NewBullet(bX, bY, 1, 200, 10, p.direction, p, angel, p.debug, p.game)
			// p.game.Screen().AddEntity(bullet)
			p.game.Screen().Level().AddEntity(bullet)
		}
	}
}
