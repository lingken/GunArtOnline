package object

import (
	"GunArtOnline/message"
	"GunArtOnline/util"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

type Enemy struct {
	Actor
	hatred int
	// target []*Actor
	target interface{}
}

func NewEnemy(name string, HP, MP, speed, posX, posY, hatred int, game *tl.Game, debug *message.DebugInfo) *Enemy {
	enemy := Enemy{
		Actor:  *NewActor(name, HP, MP, speed, posX, posY, game, debug),
		hatred: hatred,
		// target: make([]*Actor, 0),
	}
	enemy.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Ch: 'E'})
	return &enemy
}

func (enemy *Enemy) Hit(bullet *Bullet) {
	if enemy.target == bullet.source {
		enemy.hatred += bullet.Damage
	} else {
		enemy.hatred -= bullet.Damage
		if enemy.hatred < 0 {
			enemy.hatred = -enemy.hatred
			enemy.target = bullet.source
		} else if enemy.hatred == 0 {
			enemy.target = nil
		}
	}
	enemy.HP -= bullet.Damage
	if v, ok := bullet.source.(*Player); ok {
		enemy.debug.AddInfo(fmt.Sprintf("Enemy Hit remain HP: %d by %s\n", enemy.HP, v.name))
	} else {
		enemy.debug.AddInfo(fmt.Sprintf("Enemy Hit remain HP: %d\n", enemy.HP))
	}
	if enemy.HP <= 0 {
		enemy.game.Screen().Level().RemoveEntity(enemy)
	}
}

func (enemy *Enemy) Draw(s *tl.Screen) {
	if enemy.speed == 0 || enemy.frame < util.Timeconst/enemy.speed {
		enemy.frame += 1
		enemy.entity.Draw(s)
		return
	}
	enemy.prevX, enemy.prevY = enemy.entity.Position()
	prevX, prevY := enemy.prevX, enemy.prevY

	if enemy.target == nil {
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
	} else {
		if v, ok := enemy.target.(*Player); ok {
			targetX, targetY := v.Position()
			newX, newY := prevX, prevY
			if prevX < targetX {
				newX = prevX + 1
			} else if prevX > targetX {
				newX = prevX - 1
			}
			if prevY < targetY {
				newY = prevY + 1
			} else if prevY > targetY {
				newY = prevY - 1
			}
			enemy.entity.SetPosition(newX, newY)
		}
	}

	enemy.frame = 0
}
