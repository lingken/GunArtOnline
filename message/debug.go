package message

import (
	tl "github.com/JoelOtter/termloop"
	// "strconv"
	// "time"
)

type DebugInfo struct {
	lines []*tl.Text
}

func NewDebugInfo() *DebugInfo {
	lineNum := 10
	debugInfo := DebugInfo{
		lines: make([]*tl.Text, lineNum),
	}
	debugInfo.lines[0] = tl.NewText(0, 0, "Debug:", tl.ColorBlack, tl.ColorWhite)
	for i := 1; i < lineNum; i++ {
		debugInfo.lines[i] = tl.NewText(0, 0, "", tl.ColorBlack, tl.ColorWhite)
	}
	return &debugInfo
}

func (debugInfo *DebugInfo) Draw(screen *tl.Screen) {
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
}
