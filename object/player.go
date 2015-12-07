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

func NewPlayer(name string, HP, MP, speed, posX, posY int, game *tl.Game, debug *message.DebugInfo, db *util.Database, reg *util.RegisterList) *Player {
	p := Player{
		Actor: *NewActor(name, HP, MP, speed, posX, posY, game, debug, db, reg),
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
		// write operation; update database
		p.db.Put(p.Key, *p)

		if p.HP <= 0 {
			p.state = actorDead
			p.game.Screen().Level().RemoveEntity(p)
			// remove entity from level; remove from local database
			p.db.Remove(p.Key)
		}
	}
}

func (p *Player) Draw(screen *tl.Screen) {
	// read operation; fetch from database
	tmp, _ := p.db.GetValue(p.Key)
	val, _ := tmp.(Player)
	*p = val
	p.entity.Draw(screen)
	// actor.entity.Draw(screen)
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
			bullet := NewBullet(bX, bY, 1, 1000, 10, p.direction, p, angel, p.debug, p.game)
			// TODO add bullet to database!!
			p.game.Screen().Level().AddEntity(bullet)
		}
		// write operation; update database
		p.db.Put(p.Key, *p)
	}
}
