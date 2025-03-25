package main

import (
	"fmt"
	"strconv"
	"syscall"
	"time"

	"github.com/atotto/clipboard"
)

var (
	user32      = syscall.NewLazyDLL("user32.dll")
	keybd_event = user32.NewProc("keybd_event")
)

const (
	VK_TAB          = 0x09
	VK_ALT          = 0x12
	VK_NUM0         = 0x60
	VK_NUM9         = 0x69
	KEYEVENTF_KEYUP = 0x02
)

func main() {
	// Obtener texto del portapapeles
	texto, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el portapapeles:", err)
		return
	}

	if texto == "" {
		fmt.Println("El portapapeles está vacío.")
		return
	}

	fmt.Println("Texto obtenido del portapapeles:", texto)

	// Simula Alt+Tab para cambiar de ventana
	sendKeyEvent(VK_ALT, false)
	sendKeyEvent(VK_TAB, false)
	sendKeyEvent(VK_TAB, true)
	sendKeyEvent(VK_ALT, true)
	time.Sleep(100 * time.Millisecond)

	// Escribe cada carácter del portapapeles con una pausa
	for _, char := range texto {
		if err := escribirCaracter(char); err != nil {
			fmt.Printf("No se pudo escribir '%c': %v\n", char, err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func escribirCaracter(char rune) error {
	fmt.Println("Código ASCII:", int(char)) // Para depuración
	asciiStr := strconv.Itoa(int(char))

	sendKeyEvent(VK_ALT, false) // Mantener ALT presionado
	time.Sleep(50 * time.Millisecond)

	for _, digit := range asciiStr {
		vk := getVKFromDigit(digit)
		if vk == 0 {
			return fmt.Errorf("dígito inválido: %c", digit)
		}
		sendKeyEvent(vk, false)
		sendKeyEvent(vk, true)
		time.Sleep(50 * time.Millisecond)
	}

	sendKeyEvent(VK_ALT, true) // Soltar ALT
	time.Sleep(50 * time.Millisecond)
	return nil
}

func getVKFromDigit(digit rune) byte {
	switch digit {
	case '0':
		return VK_NUM0
	case '1':
		return VK_NUM0 + 1
	case '2':
		return VK_NUM0 + 2
	case '3':
		return VK_NUM0 + 3
	case '4':
		return VK_NUM0 + 4
	case '5':
		return VK_NUM0 + 5
	case '6':
		return VK_NUM0 + 6
	case '7':
		return VK_NUM0 + 7
	case '8':
		return VK_NUM0 + 8
	case '9':
		return VK_NUM0 + 9
	}
	return 0
}

func sendKeyEvent(key byte, keyUp bool) {
	var bVk, bScan byte
	var dwFlags, dwExtraInfo uint32

	bVk = key
	if keyUp {
		dwFlags = KEYEVENTF_KEYUP
	}

	keybd_event.Call(uintptr(bVk), uintptr(bScan), uintptr(dwFlags), uintptr(dwExtraInfo))
}