package inputevent

import (
	wa "github.com/ilinovalex86/winapi"
	"time"
)

func InitIE() {
	return
}

func EventsRun(events []Event) {
	m := wa.MouseEvent{}
	k := wa.KeyboardEvent{}
	for _, e := range events {
		switch e.Method {
		case "mouse":
			switch e.Event {
			case "move":
				m.Move(e.CorX, e.CorY)
			case "left":
				if e.Shift {
					k.ShiftPress()
				}
				if e.Ctrl {
					k.CtrlPress()
				}
				time.Sleep(time.Millisecond)
				m.LClick(e.CorX, e.CorY)
				time.Sleep(time.Millisecond)
				if e.Shift {
					k.ShiftRelease()
				}
				if e.Ctrl {
					k.CtrlRelease()
				}
			case "right":
				m.RClick(e.CorX, e.CorY)
			case "dbclick":
				m.DoubleClick(e.CorX, e.CorY)
			case "drop":
				m.Drop(e.CorX, e.CorY)
			case "scrollUp":
				m.WheelUp()
			case "scrollDown":
				m.WheelDown()
			}
		case "keyboard":
			k = wa.KeyboardEvent{Ctrl: e.Ctrl, Shift: e.Shift, JavaScriptCode: e.Code}
			k.Launching()
		}
		time.Sleep(time.Millisecond)
	}
}
