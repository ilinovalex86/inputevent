package inputevent

import "time"

func (e *Event) mouseLeft() {
	if e.Shift {
		shiftPress()
		defer shiftRelease()
	}
	if e.Ctrl {
		ctrlPress()
		defer ctrlRelease()
	}
	e.mouseMove()
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0002), 0, 0, 0, 0)
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0004), 0, 0, 0, 0)
}

func (e *Event) mouseRight() {
	e.mouseMove()
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0008), 0, 0, 0, 0)
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0010), 0, 0, 0, 0)
}

func (e *Event) mouseDbClick() {
	e.mouseMove()
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0002), 0, 0, 0, 0)
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0004), 0, 0, 0, 0)
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0002), 0, 0, 0, 0)
	time.Sleep(time.Millisecond)
	procMouseEvent.Call(uintptr(0x0004), 0, 0, 0, 0)
}

func (e *Event) mouseMove() {
	procSetCursorPos.Call(uintptr(e.CorX), uintptr(e.CorY))
}

func (e *Event) mouseDrop() {
	procMouseEvent.Call(uintptr(0x0002), 0, 0, 0, 0)
	time.Sleep(5 * time.Millisecond)
	procSetCursorPos.Call(uintptr(e.CorX), uintptr(e.CorY))
	time.Sleep(25 * time.Millisecond)
	procMouseEvent.Call(uintptr(0x0004), 0, 0, 0, 0)
}
