package inputevent

import (
	"time"
)

func shiftPress() {
	downKey(_VK_SHIFT)
}

func shiftRelease() {
	upKey(_VK_SHIFT)
}

func ctrlPress() {
	downKey(_VK_CTRL)
}

func ctrlRelease() {
	upKey(_VK_CTRL)
}

func (e *Event) keyboard() {
	if key, ok := codeToInt[e.Code]; ok {
		if e.Shift {
			shiftPress()
			defer shiftRelease()
		}
		if e.Ctrl {
			ctrlPress()
			defer ctrlRelease()
		}
		downKey(key)
		time.Sleep(time.Millisecond)
		upKey(key)
	}
}

func downKey(key int) {
	flag := 0
	if key < 0xFFF { // Detect if the key code is virtual or no
		flag |= _KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}

func upKey(key int) {
	flag := _KEYEVENTF_KEYUP
	if key < 0xFFF {
		flag |= _KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}

const (
	// I add 0xFFF for all Virtual key
	_VK_SHIFT           = 0x10 + 0xFFF
	_VK_CTRL            = 0x11 + 0xFFF
	_KEYEVENTF_KEYUP    = 0x0002
	_KEYEVENTF_SCANCODE = 0x0008
)

var codeToInt = map[string]int{
	"Escape": 1,

	"Insert":   0x2D + 0xFFF,
	"Delete":   0x2E + 0xFFF,
	"Home":     0x24 + 0xFFF,
	"End":      0x23 + 0xFFF,
	"PageUp":   0x21 + 0xFFF,
	"PageDown": 0x22 + 0xFFF,

	"Backquote": 41,
	"Digit1":    2,
	"Digit2":    3,
	"Digit3":    4,
	"Digit4":    5,
	"Digit5":    6,
	"Digit6":    7,
	"Digit7":    8,
	"Digit8":    9,
	"Digit9":    10,
	"Digit0":    11,
	"Minus":     12,
	"Equal":     13,
	"Backspace": 14,

	"KeyQ":         16,
	"KeyW":         17,
	"KeyE":         18,
	"KeyR":         19,
	"KeyT":         20,
	"KeyY":         21,
	"KeyU":         22,
	"KeyI":         23,
	"KeyO":         24,
	"KeyP":         25,
	"BracketLeft":  26,
	"BracketRight": 27,
	"Backslash":    43,

	"CapsLock":  58,
	"KeyA":      30,
	"KeyS":      31,
	"KeyD":      32,
	"KeyF":      33,
	"KeyG":      34,
	"KeyH":      35,
	"KeyJ":      36,
	"KeyK":      37,
	"KeyL":      38,
	"Semicolon": 39,
	"Quote":     40,
	"Enter":     28,

	"KeyZ":   44,
	"KeyX":   45,
	"KeyC":   46,
	"KeyV":   47,
	"KeyB":   48,
	"KeyN":   49,
	"KeyM":   50,
	"Comma":  51,
	"Period": 52,
	"Slash":  53,

	"Space": 57,

	"ArrowUp":    0x26 + 0xFFF,
	"ArrowDown":  0x28 + 0xFFF,
	"ArrowRight": 0x27 + 0xFFF,
	"ArrowLeft":  0x25 + 0xFFF,

	"Numpad1":        2,
	"Numpad2":        3,
	"Numpad3":        4,
	"Numpad4":        5,
	"Numpad5":        6,
	"Numpad6":        7,
	"Numpad7":        8,
	"Numpad8":        9,
	"Numpad9":        10,
	"Numpad0":        11,
	"NumpadDecimal":  83,
	"NumpadEnter":    28,
	"NumpadAdd":      78,
	"NumpadSubtract": 74,
	"NumpadMultiply": 0x6A + 0xFFF,
	"NumpadDivide":   0x6F + 0xFFF,
}
