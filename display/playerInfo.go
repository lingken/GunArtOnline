package display

import (
	"strconv"

	"GunArtOnline/object"
	tl "github.com/JoelOtter/termloop"
)

// PlayerInfo represents a player status display
type PlayerInfo struct {
	Status
	textLine1 *tl.Text
	textLine2 *tl.Text
	textLine3 *tl.Text
	textLine4 *tl.Text
	textLine5 *tl.Text
	// textLine6 *tl.Text
	// textLine7 *tl.Text
	// textLine8 *tl.Text
	player *object.Player
}

// NewPlayerStatus creates a new status display for the specified PlayerMech
func NewPlayerStatus(x, y, width, height int, player *object.Player, level *tl.BaseLevel) *PlayerInfo {
	playerStatus := PlayerInfo{
		Status: *NewStatus(x, y, width, height, level),
		player: player,
	}

	playerStatus.textLine1 = tl.NewText(x, y, "--1--", tl.ColorWhite, tl.ColorBlack)
	playerStatus.textLine2 = tl.NewText(x, y, "--2--", tl.ColorWhite, tl.ColorBlack)
	playerStatus.textLine3 = tl.NewText(x, y, "--3--", tl.ColorWhite, tl.ColorBlack)
	playerStatus.textLine4 = tl.NewText(x, y, "--4--", tl.ColorWhite, tl.ColorBlack)
	playerStatus.textLine5 = tl.NewText(x, y, "--5--", tl.ColorWhite, tl.ColorBlack)
	// playerStatus.textLine6 = tl.NewText(x, y, "--6--", tl.ColorWhite, tl.ColorBlack)
	// playerStatus.textLine7 = tl.NewText(x, y, "--7--", tl.ColorWhite, tl.ColorBlack)
	// playerStatus.textLine8 = tl.NewText(x, y, "--8--", tl.ColorWhite, tl.ColorBlack)

	return &playerStatus
}

// Draw passes the draw call to entity.
func (display *PlayerInfo) Draw(screen *tl.Screen) {

	offSetX, offSetY := display.level.Offset()

	display.background.SetPosition(-offSetX+display.x, -offSetY+display.y)
	display.textLine1.SetPosition(-offSetX+1+display.x, -offSetY+1+display.y)
	display.textLine2.SetPosition(-offSetX+1+display.x, -offSetY+3+display.y)
	display.textLine3.SetPosition(-offSetX+1+display.x, -offSetY+5+display.y)
	display.textLine4.SetPosition(-offSetX+1+display.x, -offSetY+7+display.y)
	display.textLine5.SetPosition(-offSetX+1+display.x, -offSetY+8+display.y)
	// display.textLine6.SetPosition(-offSetX+1+display.x, -offSetY+9+display.y)
	// display.textLine7.SetPosition(-offSetX+1+display.x, -offSetY+10+display.y)
	// display.textLine8.SetPosition(-offSetX+1+display.x, -offSetY+11+display.y)

	display.background.Draw(screen)
	display.textLine1.Draw(screen)
	display.textLine2.Draw(screen)
	display.textLine3.Draw(screen)
	display.textLine4.Draw(screen)
	display.textLine5.Draw(screen)
	// display.textLine6.Draw(screen)
	// display.textLine7.Draw(screen)
	// display.textLine8.Draw(screen)
}

// Tick is called to process 1 tick of actions based on the
// type of event.
func (display *PlayerInfo) Tick(event tl.Event) {
	display.textLine1.SetText(display.player.name)
	display.textLine2.SetText("       HP: " + strconv.Itoa(display.player.HP))
	//x, y := display.player.Position()
	display.textLine3.SetText("       MP: " + strconv.Itoa(display.player.MP))

	display.textLine4.SetText("    speed: " + strconv.Itoa(display.player.speed))
	//assume for now there is only 1 Weapon
	display.textLine5.SetText("direction: " + strconv.Itoa(display.player.direction))
}
