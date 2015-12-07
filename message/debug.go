package message

import (
	"GunArtOnline/util"
	"encoding/json"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"log"
	"os"
	"strconv"
)

type DebugInfo struct {
	lines []*tl.Text
	db    *util.Database
	reg   *util.RegisterList
	Key   string
	frame int
}

func NewDebugInfo(db *util.Database, reg *util.RegisterList) *DebugInfo {
	lineNum := 5
	debugInfo := DebugInfo{
		lines: make([]*tl.Text, lineNum),
		db:    db,
		reg:   reg,
		Key:   "debugInfo",
	}
	debugInfo.lines[0] = tl.NewText(0, 0, "Debug: ", tl.ColorBlack, tl.ColorWhite)
	for i := 1; i < lineNum; i++ {
		debugInfo.lines[i] = tl.NewText(0, 0, "", tl.ColorBlack, tl.ColorWhite)
	}

	// debugInfo.Put(db)
	// debugInfo.db.Put(debugInfo.Key, debugInfo)

	return &debugInfo
}

func (debugInfo *DebugInfo) Draw(screen *tl.Screen) {

	_ = debugInfo.GetValue(debugInfo.db)

	// val, _ := debugInfo.db.GetValue(debugInfo.Key)
	// tmp, _ := val.(DebugInfo) // type cast from interface{} to DebugInfo
	// *debugInfo = tmp          // update the content of debugInfo at this machine used by game engine

	_, screenHeight := screen.Size()
	for i := 0; i < len(debugInfo.lines); i++ {
		debugInfo.lines[i].SetPosition(4, screenHeight-15+i)
		debugInfo.lines[i].Draw(screen)
	}
}

func (debugInfo *DebugInfo) Tick(e tl.Event) {}

func (debugInfo *DebugInfo) AddInfo(info string) {
	f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)

	log.Println("AddInfo: ", info)
	// log everything

	for i := 1; i < len(debugInfo.lines)-1; i++ {
		debugInfo.lines[i].SetText(debugInfo.lines[i+1].Text())
		log.Println(i, debugInfo.lines[i].Text()) // log
	}
	debugInfo.lines[len(debugInfo.lines)-1].SetText(info)
	log.Println(len(debugInfo.lines)-1, debugInfo.lines[len(debugInfo.lines)-1].Text()) // log

	// write operation; write to database
	debugInfo.Put(debugInfo.db)
	// debugInfo.db.Put(debugInfo.Key, *debugInfo)
}

func (debugInfo *DebugInfo) Put(db *util.Database) {
	f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	defer f.Close()

	// assign it to the standard logger
	log.SetOutput(f)

	log.Println("Put something... Length of debugInfo.lines: ", len(debugInfo.lines))
	// log out all stuff
	// for i := 1; i < len(debugInfo.lines); i++ {
	// 	log.Println(i, debugInfo.lines[i].Text())
	// 	// data, _ := json.Marshal((*debugInfo.lines[i]))
	// 	// log.Println(i, data)
	// }

	for i := 1; i < len(debugInfo.lines); i++ {
		// log.Println(i, debugInfo.lines[i].Text()) // log

		key := util.GenerateKey(debugInfo.Key, "line") + "_" + strconv.Itoa(i)

		data, err := json.Marshal(debugInfo.lines[i].Text())
		if err != nil {
			debugInfo.AddInfo("debugInfo Put Error")
		}

		log.Println(i, key, data)
		db.PaxosPut(key, data)
	}
}

func (debugInfo *DebugInfo) GetValue(db *util.Database) bool {
	// f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
	// }

	// // don't forget to close it
	// defer f.Close()

	// // assign it to the standard logger
	// log.SetOutput(f)

	// log.Println("Get something...")
	for i := 1; i < len(debugInfo.lines); i++ {
		key := util.GenerateKey(debugInfo.Key, "line") + "_" + strconv.Itoa(i)

		// var data tl.Text
		var str string
		if v, success := db.PaxosGetValue(key); !success {
			debugInfo.AddInfo("debugInfo GetValue Error")
			return false
		} else {
			if value, ok := v.([]byte); ok {
				json.Unmarshal(value, &str)
				debugInfo.lines[i].SetText(str)
			} else {
				debugInfo.AddInfo("debugInfo GetValue Error Unmarshal")
				return false
			}
			// Log info
			// log.Println(debugInfo.lines[i].Text())
		}
	}

	return true
}
