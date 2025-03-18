package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/micmonay/keybd_event"
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

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// Simula Alt+Tab
	kb.SetKeys(keybd_event.VK_TAB)
	kb.HasALT(true)
	kb.Press()
	kb.Release()

	time.Sleep(250 * time.Millisecond) // Espera para cambiar de ventana

	// Escribe cada carácter del portapapeles con una pausa
	for _, char := range texto {
		if err := EscribirCaracter(kb, char); err != nil {
			fmt.Printf("No se pudo escribir '%c': %v\n", char, err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// EscribirCaracter envía un carácter al teclado virtual
func EscribirCaracter(kb keybd_event.KeyBonding, char rune) error {
	kb.Clear() // Limpia cualquier tecla anterior

	shift := false
	key, shiftRequired := ConvertirCaracter(char)

	// Si el carácter requiere Shift, activamos Shift
	if shiftRequired {
		kb.HasSHIFT(true)
		shift = true
	}

	if key != -1 {
		kb.SetKeys(key)
		kb.Press()
		kb.Release()
	}

	// Si activamos Shift, lo desactivamos
	if shift {
		kb.HasSHIFT(false)
	}

	return nil
}

// ConvertirCaracter mapea caracteres a teclas del teclado virtual
func ConvertirCaracter(char rune) (int, bool) {
	switch char {
	case 'a':
		return keybd_event.VK_A, false
	case 'b':
		return keybd_event.VK_B, false
	case 'c':
		return keybd_event.VK_C, false
	case 'd':
		return keybd_event.VK_D, false
	case 'e':
		return keybd_event.VK_E, false
	case 'f':
		return keybd_event.VK_F, false
	case 'g':
		return keybd_event.VK_G, false
	case 'h':
		return keybd_event.VK_H, false
	case 'i':
		return keybd_event.VK_I, false
	case 'j':
		return keybd_event.VK_J, false
	case 'k':
		return keybd_event.VK_K, false
	case 'l':
		return keybd_event.VK_L, false
	case 'm':
		return keybd_event.VK_M, false
	case 'n':
		return keybd_event.VK_N, false
	case 'o':
		return keybd_event.VK_O, false
	case 'p':
		return keybd_event.VK_P, false
	case 'q':
		return keybd_event.VK_Q, false
	case 'r':
		return keybd_event.VK_R, false
	case 's':
		return keybd_event.VK_S, false
	case 't':
		return keybd_event.VK_T, false
	case 'u':
		return keybd_event.VK_U, false
	case 'v':
		return keybd_event.VK_V, false
	case 'w':
		return keybd_event.VK_W, false
	case 'x':
		return keybd_event.VK_X, false
	case 'y':
		return keybd_event.VK_Y, false
	case 'z':
		return keybd_event.VK_Z, false
	case 'A':
		return keybd_event.VK_A, true
	case 'B':
		return keybd_event.VK_B, true
	case 'C':
		return keybd_event.VK_C, true
	case 'D':
		return keybd_event.VK_D, true
	case 'E':
		return keybd_event.VK_E, true
	case 'F':
		return keybd_event.VK_F, true
	case 'G':
		return keybd_event.VK_G, true
	case 'H':
		return keybd_event.VK_H, true
	case 'I':
		return keybd_event.VK_I, true
	case 'J':
		return keybd_event.VK_J, true
	case 'K':
		return keybd_event.VK_K, true
	case 'L':
		return keybd_event.VK_L, true
	case 'M':
		return keybd_event.VK_M, true
	case 'N':
		return keybd_event.VK_N, true
	case 'O':
		return keybd_event.VK_O, true
	case 'P':
		return keybd_event.VK_P, true
	case 'Q':
		return keybd_event.VK_Q, true
	case 'R':
		return keybd_event.VK_R, true
	case 'S':
		return keybd_event.VK_S, true
	case 'T':
		return keybd_event.VK_T, true
	case 'U':
		return keybd_event.VK_U, true
	case 'V':
		return keybd_event.VK_V, true
	case 'W':
		return keybd_event.VK_W, true
	case 'X':
		return keybd_event.VK_X, true
	case 'Y':
		return keybd_event.VK_Y, true
	case 'Z':
		return keybd_event.VK_Z, true
	case '0':
		return keybd_event.VK_0, false
	case '1':
		return keybd_event.VK_1, false
	case '2':
		return keybd_event.VK_2, false
	case '3':
		return keybd_event.VK_3, false
	case '4':
		return keybd_event.VK_4, false
	case '5':
		return keybd_event.VK_5, false
	case '6':
		return keybd_event.VK_6, false
	case '7':
		return keybd_event.VK_7, false
	case '8':
		return keybd_event.VK_8, false
	case '9':
		return keybd_event.VK_9, false
	case ' ':
		return keybd_event.VK_SPACE, false
	case '.':
		return keybd_event.VK_DOT, false
	case ',':
		return keybd_event.VK_COMMA, false
	case ';':
		return keybd_event.VK_SEMICOLON, false
	case ':':
		return keybd_event.VK_SEMICOLON, true
	case '-':
		return keybd_event.VK_MINUS, false
	case '_':
		return keybd_event.VK_MINUS, true
	case '=':
		return keybd_event.VK_EQUAL, false
	case '+':
		return keybd_event.VK_EQUAL, true
	case '(':
		return keybd_event.VK_9, true
	case ')':
		return keybd_event.VK_0, true
	case '!':
		return keybd_event.VK_1, true
	case '?':
		return keybd_event.VK_SLASH, true
	case '/':
		return keybd_event.VK_SLASH, false
	default:
		return -1, false // Caracter no soportado
	}
}
