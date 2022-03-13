package inputevent

import (
	"syscall"
	"time"
)

var (
	moduser32        = syscall.NewLazyDLL("user32.dll")
	procMouseEvent   = moduser32.NewProc("mouse_event")
	procSetCursorPos = moduser32.NewProc("SetCursorPos")
	procKeyBd        = moduser32.NewProc("keybd_event")
)

func InitIE() {
	return
}

func EventsRun(events []Event) {
	for _, event := range events {
		e := &event
		//fmt.Printf("%#v \n", *e)
		switch e.Method {
		case "mouse":
			switch e.Event {
			case "move":
				e.mouseMove()
			case "left":
				e.mouseLeft()
			case "right":
				e.mouseRight()
			case "dbclick":
				e.mouseDbClick()
			case "drop":
				e.mouseDrop()
			}
		case "keyboard":
			e.keyboard()
		}
		time.Sleep(time.Millisecond)
	}
}
