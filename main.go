package main

import (
	"fmt"
	"unicode"
	//"strconv"
	"time"
	"syscall"
	"github.com/atotto/clipboard"
)

const debugging = false

func main() {
	// Obtener texto del portapapeles
	texto, err := clipboard.ReadAll()
	if err != nil {
		if debugging {
			fmt.Println("Error al leer el portapapeles:", err)
		}
		return
	}

	if texto == "" {
		if debugging {
			fmt.Println("El portapapeles está vacío.")
		}
		return
	}

	if debugging {
		fmt.Println("Texto obtenido del portapapeles:", texto)
	}

	kb, err := NewKeyBonding()
	if err != nil {
		if debugging {
			fmt.Println("Error al inicializar el teclado virtual:", err)
		}
		return
	}

	// Simula Alt+Tab para cambiar de ventana
	kb.SetKeys(VK_TAB)
	kb.HasALT(true)
	kb.Press()
	kb.Release()
	time.Sleep(100 * time.Millisecond) // Espera para cambiar de ventana

	// Escribe cada carácter del portapapeles con una pausa
	for _, char := range texto {
		if err := escribirCaracter(kb, char); err != nil {
			if debugging {
				fmt.Printf("No se pudo escribir '%c': %v\n", char, err)
			}
		}
	}
}

func getAltCode(c rune) string {
	if unicode.IsPrint(c) {
		// Convertir a string con ceros a la izquierda (para que siempre tenga 4 dígitos)
		return fmt.Sprintf("%04d", int(c))
	}
	return "Carácter no imprimible"
}

// escribirCaracter envía un carácter al teclado virtual
func escribirCaracter(kb KeyBonding, char rune) error {
	kb.Clear()
	altCode := getAltCode(char)
	if debugging {
		fmt.Println("Código ASCII:", altCode) // Para depuración
	}

	// Convertimos el código ASCII en string
	//asciiStr := strconv.Itoa(int(char))
	asciiStr := altCode

	// Mantener ALT presionado
	kb.HasALT(true)
	kb.Press()
	time.Sleep(10 * time.Millisecond)

	// Convertir cada dígito del código ASCII en su respectiva tecla
	var keys []int
	for _, digit := range asciiStr {
		vk := getVKFromDigit(digit)
		if vk == -1 {
			if debugging {
				fmt.Printf("No se puede escribir el dígito '%c'\n", digit)
			}
			continue
		}
		keys = append(keys, vk) // Asegura que cada dígito se añade sin perder ceros
	}

	// Presionar cada número del código ASCII
	if len(keys) > 0 {
		if debugging {
			fmt.Printf("Teclas a presionar: %v\n", keys)
		}
		kb.SetKeys(keys...)
		kb.Press()
	}

	// Soltar ALT
	kb.Release()
	time.Sleep(10 * time.Millisecond)
	return nil
}

func getVKFromDigit(digit rune) int {
	switch digit {
	case '0':
		return VK_KP0
	case '1':
		return VK_KP1
	case '2':
		return VK_KP2
	case '3':
		return VK_KP3
	case '4':
		return VK_KP4
	case '5':
		return VK_KP5
	case '6':
		return VK_KP6
	case '7':
		return VK_KP7
	case '8':
		return VK_KP8
	case '9':
		return VK_KP9
	}
	return -1
}


var dll = syscall.NewLazyDLL("user32.dll")
var procKeyBd = dll.NewProc("keybd_event")

// Press key(s)
func (k *KeyBonding) Press() error {
	if k.hasALT {
		downKey(_VK_ALT)
	}
	for _, key := range k.keys {
		downKey(key)
		upKey(key)
	}
	return nil
}

//Release key(s)
func (k *KeyBonding) Release() error {
	if k.hasALT {
		upKey(_VK_ALT)
	}
	for _, key := range k.keys {
		upKey(key)
	}
	return nil
}

// Launch key bounding
func (k *KeyBonding) Launching() error {
	err := k.Press()
	if err != nil {
		return err
	}
	err = k.Release()
	return err
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
func initKeyBD() error { return nil }

//KeyBonding type for keybd_event
type KeyBonding struct {
	hasCTRL   bool
	hasALT    bool
	hasSHIFT  bool
	hasRCTRL  bool
	hasRSHIFT bool
	hasALTGR  bool
	hasSuper  bool
	keys      []int
}

//NewKeyBonding Use for create struct KeyBounding
func NewKeyBonding() (KeyBonding, error) {
	keyBounding := KeyBonding{}
	keyBounding.Clear()
	err := initKeyBD()
	if err != nil {
		return keyBounding, err
	}
	return keyBounding, nil
}

//Clear clean instance
func (k *KeyBonding) Clear() {
	k.hasALT = false
	k.hasCTRL = false
	k.hasSHIFT = false
	k.hasRCTRL = false
	k.hasRSHIFT = false
	k.hasALTGR = false
	k.hasSuper = false
	k.keys = []int{}
}

//SetKeys set keys
func (k *KeyBonding) SetKeys(keys ...int) {
	k.keys = keys
}

//AddKey add one key pressed
func (k *KeyBonding) AddKey(key int) {
	k.keys = append(k.keys, key)
}

//HasALT If key ALT pressed
func (k *KeyBonding) HasALT(b bool) {
	k.hasALT = b
}

const (
	_VK_ALT             = 0x12 + 0xFFF
	_KEYEVENTF_KEYUP    = 0x0002
	_KEYEVENTF_SCANCODE = 0x0008
	VK_TAB     = 15
	VK_KP0     = 82
	VK_KP1     = 79
	VK_KP2     = 80
	VK_KP3     = 81
	VK_KP4     = 75
	VK_KP5     = 76
	VK_KP6     = 77
	VK_KP7     = 71
	VK_KP8     = 72
	VK_KP9     = 73
)

