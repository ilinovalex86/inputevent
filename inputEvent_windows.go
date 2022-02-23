package inputevent

import "fmt"

func EventsRun(events []Event) {
	for _, e := range events {
		fmt.Printf("%#v \n", e)
	}
}

//func EventsRun(events []cn.Event) {
//	for _, e := range events {
//		fmt.Printf("%#v \n", e)
//		switch e.Method {
//		case "mouse":
//			switch e.Event {
//			case "move":
//				robotgo.Move(e.CorX, e.CorY)
//			case "left":
//				if e.Shift {
//					robotgo.KeyToggle("shift")
//				}
//				if e.Ctrl {
//					robotgo.KeyToggle("ctrl")
//				}
//				robotgo.MoveClick(e.CorX, e.CorY)
//				if e.Shift {
//					robotgo.KeyToggle("shift", "up")
//				}
//				if e.Ctrl {
//					robotgo.KeyToggle("ctrl", "up")
//				}
//			case "right":
//				robotgo.MoveClick(e.CorX, e.CorY, "right")
//			case "dbclick":
//				robotgo.MoveClick(e.CorX, e.CorY, "left", true)
//			case "drop":
//				robotgo.DragSmooth(e.CorX, e.CorY)
//			}
//		case "keyboard":
//			if key, ok := keys[e.Event]; ok {
//				if e.Shift && e.Ctrl {
//					robotgo.KeyTap(key, "shift", "ctrl")
//				} else if e.Shift {
//					robotgo.KeyTap(key, "shift")
//				} else if e.Ctrl {
//					robotgo.KeyTap(key, "ctrl")
//				} else {
//					robotgo.KeyTap(key)
//				}
//				continue
//			}
//			robotgo.TypeStr(e.Event)
//		}
//		time.Sleep(10 * time.Millisecond)
//	}
//}
