package weapon

import (
	"GunArtOnline/util"
)

type Bullet struct {
	posX      int
	posY      int
	symbol    rune
	damage    int
	speed     int
	direction util.Direction
}
