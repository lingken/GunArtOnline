package message

import (
	"GunArtOnline/util"
	tl "github.com/JoelOtter/termloop"
	// "strconv"
	// "time"
)

type DebugInfo struct {
	lines []*tl.Text
	db    *util.Database
	reg   *util.RegisterList
	Key   string
}

func NewDebugInfo(db *util.Database, reg *util.RegisterList) *DebugInfo {
	lineNum := 10
	debugInfo := DebugInfo{
		lines: make([]*tl.Text, lineNum),
		db:    db,
		reg:   reg,
		Key:   "debugInfo",
	}
	debugInfo.lines[0] = tl.NewText(0, 0, "Debug:", tl.ColorBlack, tl.ColorWhite)
	for i := 1; i < lineNum; i++ {
		debugInfo.lines[i] = tl.NewText(0, 0, "", tl.ColorBlack, tl.ColorWhite)
	}
	return &debugInfo
}

func (debugInfo *DebugInfo) Draw(screen *tl.Screen) {
	// read operation; fetch from database
	val, _ := debugInfo.db.GetValue(debugInfo.Key)
	tmp, _ := val.(DebugInfo) // type cast from interface{} to DebugInfo
	*debugInfo = tmp          // update the content of debugInfo at this machine used by game engine

	_, screenHeight := screen.Size()
	for i := 0; i < len(debugInfo.lines); i++ {
		debugInfo.lines[i].SetPosition(4, screenHeight-15+i)
		debugInfo.lines[i].Draw(screen)
	}
}

func (debugInfo *DebugInfo) Tick(e tl.Event) {}

func (debugInfo *DebugInfo) AddInfo(info string) {
	for i := 1; i < len(debugInfo.lines)-1; i++ {
		debugInfo.lines[i].SetText(debugInfo.lines[i+1].Text())
	}
	debugInfo.lines[len(debugInfo.lines)-1].SetText(info)

	// write operation; write to database
	debugInfo.db.Put(debugInfo.Key, *debugInfo)
}
