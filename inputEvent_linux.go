package inputevent

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

var keys = map[string]string{
	"Home":       "home",
	"End":        "end",
	"PageUp":     "pageup",
	"PageDown":   "pagedown",
	"ArrowUp":    "up",
	"ArrowDown":  "down",
	"ArrowRight": "right",
	"ArrowLeft":  "left",
	"Backspace":  "backspace",
	"Delete":     "delete",
	"Enter":      "enter",
	"Escape":     "escape",
	"Insert":     "insert",
	"Space":      "space",
	"Tab":        "tab",
	"KeyZ":       "z",
	"KeyX":       "x",
	"KeyC":       "c",
	"KeyV":       "v",
	"KeyA":       "a",
}

func InitIE() {
	robotgo.KeyToggle("shift")
	robotgo.KeyToggle("shift", "up")
}

func EventsRun(events []Event) {
	for _, e := range events {
		fmt.Printf("%#v \n", e)
		switch e.Method {
		case "mouse":
			switch e.Event {
			case "move":
				robotgo.Move(e.CorX, e.CorY)
			case "left":
				if e.Shift {
					robotgo.KeyToggle("shift")
				}
				if e.Ctrl {
					robotgo.KeyToggle("ctrl")
				}
				robotgo.MoveClick(e.CorX, e.CorY)
				if e.Shift {
					robotgo.KeyToggle("shift", "up")
				}
				if e.Ctrl {
					robotgo.KeyToggle("ctrl", "up")
				}
			case "right":
				robotgo.MoveClick(e.CorX, e.CorY, "right")
			case "dbclick":
				robotgo.MoveClick(e.CorX, e.CorY, "left", true)
			case "drop":
				robotgo.DragSmooth(e.CorX, e.CorY)
			}
		case "keyboard":
			if key, ok := keys[e.Key]; ok {
				if e.Shift && e.Ctrl {
					robotgo.KeyTap(key, "shift", "ctrl")
				} else if e.Shift {
					robotgo.KeyTap(key, "shift")
				} else if e.Ctrl {
					robotgo.KeyTap(key, "ctrl")
				} else {
					robotgo.KeyTap(key)
				}
				continue
			}
			robotgo.TypeStr(e.Key)
		}
		time.Sleep(2 * time.Millisecond)
	}
}
