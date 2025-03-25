package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
	"unicode"
	"unsafe"
)

const (
	debugging = true
	
	// Virtual-Key Codes
	VK_ALT             = 0x12 + 0xFFF
	KEYEVENTF_KEYUP    = 0x0002
	KEYEVENTF_SCANCODE = 0x0008
	VK_TAB     = 15

	// Clipboard format for Unicode text
	CF_UNICODETEXT = 13
)

var (
	user32                     = syscall.MustLoadDLL("user32")
	isClipboardFormatAvailable = user32.MustFindProc("IsClipboardFormatAvailable")
	openClipboard              = user32.MustFindProc("OpenClipboard")
	closeClipboard             = user32.MustFindProc("CloseClipboard")
	getClipboardData           = user32.MustFindProc("GetClipboardData")
	procKeyBd                  = user32.MustFindProc("keybd_event")

	kernel32     = syscall.MustLoadDLL("kernel32")
	globalLock   = kernel32.MustFindProc("GlobalLock")
	globalUnlock = kernel32.MustFindProc("GlobalUnlock")

	// Numeric Keypad Virtual Keys
	VK_NUMPAD = [...]int{82, 79, 80, 81, 75, 76, 77, 71, 72, 73} // 0-9
)

type KeyBonding struct {
	keys []int
}

func main() {
	// Retrieve clipboard text
	text, err := ReadAll()
	if err != nil || text == "" {
		if debugging {
			fmt.Println("Clipboard read error or empty clipboard:", err)
		}
		return
	}

	if debugging {
		fmt.Println("Clipboard text:", text)
	}

	kb, err := NewKeyBonding()
	if err != nil {
		if debugging {
			fmt.Println("Virtual keyboard initialization error:", err)
		}
		return
	}

	// Simulate Alt+Tab to switch windows
	kb.SetKeys(VK_TAB)
	kb.Press()
	time.Sleep(100 * time.Millisecond)

	// Type clipboard text character by character
	for _, char := range text {
		if err := typeCharacter(kb, char); err != nil && debugging {
			fmt.Printf("Failed to type '%c': %v\n", char, err)
		}
	}
}

// Converts a rune into a four-digit Alt code string
func getAltCode(c rune) string {
	if int(c) == 10 {
		return "0010"
	} else if unicode.IsPrint(c) {
		return fmt.Sprintf("%04d", int(c))
	}
	return "0000" // Non-printable characters are ignored
}

// Simulates typing a character using Alt codes
func typeCharacter(kb KeyBonding, char rune) error {
	altCode := getAltCode(char)
	if debugging {
		fmt.Println("UNICODE Code:", altCode)
	}

	var keys []int
	for _, digit := range altCode {
		vk := getVKFromDigit(digit)
		if vk != -1 {
			keys = append(keys, vk)
		}
	}

	if len(keys) > 0 {
		if debugging {
			fmt.Printf("Keys to press: %v\n", keys)
		}
		kb.SetKeys(keys...)
		kb.Press()
	}

	time.Sleep(10 * time.Millisecond)
	return nil
}

// Maps digit runes to Virtual-Key Codes
func getVKFromDigit(digit rune) int {
	if digit >= '0' && digit <= '9' {
		return VK_NUMPAD[digit-'0']
	}
	return -1
}

// KeyBonding Methods
func (k *KeyBonding) Press() error {
	downKey(VK_ALT)
	for _, key := range k.keys {
		downKey(key)
		upKey(key)
	}
	upKey(VK_ALT)
	return nil
}

func upKey(key int) {
	flag := KEYEVENTF_KEYUP
	if key < 0xFFF {
		flag |= KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}

func downKey(key int) {
	flag := 0
	if key < 0xFFF {
		flag |= KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
}

func NewKeyBonding() (KeyBonding, error) {
	return KeyBonding{}, nil
}

func (k *KeyBonding) SetKeys(keys ...int) {
	k.keys = keys
}

// Clipboard Functions
func ReadAll() (string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if available, _, err := isClipboardFormatAvailable.Call(CF_UNICODETEXT); available == 0 {
		return "", err
	}

	if err := waitOpenClipboard(); err != nil {
		return "", err
	}

	h, _, err := getClipboardData.Call(CF_UNICODETEXT)
	if h == 0 {
		closeClipboard.Call()
		return "", err
	}

	l, _, err := globalLock.Call(h)
	if l == 0 {
		closeClipboard.Call()
		return "", err
	}

	text := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(l))[:])
	globalUnlock.Call(h)
	closeClipboard.Call()

	return text, nil
}

func waitOpenClipboard() error {
	start := time.Now()
	limit := start.Add(time.Second)

	for time.Now().Before(limit) {
		if r, _, _ := openClipboard.Call(0); r != 0 {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	return fmt.Errorf("failed to open clipboard")
}
