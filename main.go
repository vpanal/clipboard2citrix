package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/micmonay/keybd_event"
)

const version  = "1.1.2 es-caps"

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
		fmt.Println("Error al inicializar el teclado virtual:", err)
		return
	}

	// Simula Alt+Tab para cambiar de ventana
	kb.SetKeys(keybd_event.VK_TAB)
	kb.HasALT(true)
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
	time.Sleep(100 * time.Millisecond) // Espera para cambiar de ventana

	// Escribe cada carácter del portapapeles con una pausa
	for _, char := range texto {
		if err := escribirCaracter(kb, char); err != nil {
			fmt.Printf("No se pudo escribir '%c': %v\n", char, err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// escribirCaracter envía un carácter al teclado virtual
func escribirCaracter(kb keybd_event.KeyBonding, char rune) error {
	kb.Clear()
	key, shiftRequired, altGRRequired, spaceRequired, accentoRequired, caretRequired, dieresisRequired, mayuscula := convertirCaracter(char)

	if accentoRequired != 0 {
		var accento int
		if accentoRequired == 1 {
			accento = keybd_event.VK_SP7
		} else {
			accento = keybd_event.VK_SP4
		}
		kb.SetKeys(accento)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}
	
	if caretRequired {
		caret := keybd_event.VK_SP4
		kb.HasSHIFT(true)
		kb.SetKeys(caret)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
		kb.HasSHIFT(false)
	}

	if dieresisRequired {
		dieresis := keybd_event.VK_SP7
		kb.HasSHIFT(true)
		kb.SetKeys(dieresis)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
		kb.HasSHIFT(false)
	}

	if shiftRequired {
		kb.HasSHIFT(true)
	}

	if mayuscula {
		mayus := keybd_event.VK_CAPSLOCK
		kb.SetKeys(mayus)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}

	if altGRRequired {
		kb.HasALTGR(true)
	}

	if key != -1 {
		kb.SetKeys(key)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}

	if mayuscula {
		mayus := keybd_event.VK_CAPSLOCK
		kb.SetKeys(mayus)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}
	kb.HasSHIFT(false) // Desactivar Shift después de escribir
	kb.HasALTGR(false) // Desactivar AltGr después de escribir

	if spaceRequired {
		key = keybd_event.VK_SPACE
		kb.SetKeys(key)
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}
	return nil
}

// convertirCaracter mapea caracteres a teclas del teclado virtual
func convertirCaracter(char rune) (int, bool, bool, bool, int, bool, bool, bool) {
	mapeo := map[rune]struct {
		key          int
		shiftRequired bool
		altGRRequired bool
		spaceRequired bool
		accentoRequired int
		caretRequired bool
		dieresisRequired bool
		mayuscula bool
	}{
		'a': {keybd_event.VK_A, false, false, false, 0, false, false, false}, 'A': {keybd_event.VK_A, false, false, false, 0, false, false, true},
		'b': {keybd_event.VK_B, false, false, false, 0, false, false, false}, 'B': {keybd_event.VK_B, false, false, false, 0, false, false, true},
		'c': {keybd_event.VK_C, false, false, false, 0, false, false, false}, 'C': {keybd_event.VK_C, false, false, false, 0, false, false, true},
		'ç': {keybd_event.VK_SP8, false, false, false, 0, false, false, false}, 'Ç': {keybd_event.VK_SP8, false, false, false, 0, false, false, true},
		'd': {keybd_event.VK_D, false, false, false, 0, false, false, false}, 'D': {keybd_event.VK_D, false, false, false, 0, false, false, true},
		'e': {keybd_event.VK_E, false, false, false, 0, false, false, false}, 'E': {keybd_event.VK_E, false, false, false, 0, false, false, true},
		'f': {keybd_event.VK_F, false, false, false, 0, false, false, false}, 'F': {keybd_event.VK_F, false, false, false, 0, false, false, true},
		'g': {keybd_event.VK_G, false, false, false, 0, false, false, false}, 'G': {keybd_event.VK_G, false, false, false, 0, false, false, true},
		'h': {keybd_event.VK_H, false, false, false, 0, false, false, false}, 'H': {keybd_event.VK_H, false, false, false, 0, false, false, true},
		'i': {keybd_event.VK_I, false, false, false, 0, false, false, false}, 'I': {keybd_event.VK_I, false, false, false, 0, false, false, true},
		'j': {keybd_event.VK_J, false, false, false, 0, false, false, false}, 'J': {keybd_event.VK_J, false, false, false, 0, false, false, true},
		'k': {keybd_event.VK_K, false, false, false, 0, false, false, false}, 'K': {keybd_event.VK_K, false, false, false, 0, false, false, true},
		'l': {keybd_event.VK_L, false, false, false, 0, false, false, false}, 'L': {keybd_event.VK_L, false, false, false, 0, false, false, true},
		'm': {keybd_event.VK_M, false, false, false, 0, false, false, false}, 'M': {keybd_event.VK_M, false, false, false, 0, false, false, true},
		'n': {keybd_event.VK_N, false, false, false, 0, false, false, false}, 'N': {keybd_event.VK_N, false, false, false, 0, false, false, true},
		'ñ': {keybd_event.VK_SP6, false, false, false, 0, false, false, false}, 'Ñ': {keybd_event.VK_SP6, false, false, false, 0, false, false, true},
		'o': {keybd_event.VK_O, false, false, false, 0, false, false, false}, 'O': {keybd_event.VK_O, false, false, false, 0, false, false, true},
		'p': {keybd_event.VK_P, false, false, false, 0, false, false, false}, 'P': {keybd_event.VK_P, false, false, false, 0, false, false, true},
		'q': {keybd_event.VK_Q, false, false, false, 0, false, false, false}, 'Q': {keybd_event.VK_Q, false, false, false, 0, false, false, true},
		'r': {keybd_event.VK_R, false, false, false, 0, false, false, false}, 'R': {keybd_event.VK_R, false, false, false, 0, false, false, true},
		's': {keybd_event.VK_S, false, false, false, 0, false, false, false}, 'S': {keybd_event.VK_S, false, false, false, 0, false, false, true},
		't': {keybd_event.VK_T, false, false, false, 0, false, false, false}, 'T': {keybd_event.VK_T, false, false, false, 0, false, false, true},
		'u': {keybd_event.VK_U, false, false, false, 0, false, false, false}, 'U': {keybd_event.VK_U, false, false, false, 0, false, false, true},
		'v': {keybd_event.VK_V, false, false, false, 0, false, false, false}, 'V': {keybd_event.VK_V, false, false, false, 0, false, false, true},
		'w': {keybd_event.VK_W, false, false, false, 0, false, false, false}, 'W': {keybd_event.VK_W, false, false, false, 0, false, false, true},
		'x': {keybd_event.VK_X, false, false, false, 0, false, false, false}, 'X': {keybd_event.VK_X, false, false, false, 0, false, false, true},
		'y': {keybd_event.VK_Y, false, false, false, 0, false, false, false}, 'Y': {keybd_event.VK_Y, false, false, false, 0, false, false, true},
		'z': {keybd_event.VK_Z, false, false, false, 0, false, false, false}, 'Z': {keybd_event.VK_Z, false, false, false, 0, false, false, true},
		'0': {keybd_event.VK_0, false, false, false, 0, false, false, false}, '1': {keybd_event.VK_1, false, false, false, 0, false, false, false},
		'2': {keybd_event.VK_2, false, false, false, 0, false, false, false}, '3': {keybd_event.VK_3, false, false, false, 0, false, false, false},
		'4': {keybd_event.VK_4, false, false, false, 0, false, false, false}, '5': {keybd_event.VK_5, false, false, false, 0, false, false, false},
		'6': {keybd_event.VK_6, false, false, false, 0, false, false, false}, '7': {keybd_event.VK_7, false, false, false, 0, false, false, false},
		'8': {keybd_event.VK_8, false, false, false, 0, false, false, false}, '9': {keybd_event.VK_9, false, false, false, 0, false, false, false},
		' ': {keybd_event.VK_SPACE, false, false, false, 0, false, false, false}, '!': {keybd_event.VK_1, true, false, false, 0, false, false, false},
		'"': {keybd_event.VK_2, true, false, false, 0, false, false, false}, '·': {keybd_event.VK_3, true, false, false, 0, false, false, false},
		'$': {keybd_event.VK_4, true, false, false, 0, false, false, false}, '%': {keybd_event.VK_5, true, false, false, 0, false, false, false},
		'&': {keybd_event.VK_6, true, false, false, 0, false, false, false}, '/': {keybd_event.VK_7, true, false, false, 0, false, false, false},
		'(': {keybd_event.VK_8, true, false, false, 0, false, false, false}, ')': {keybd_event.VK_9, true, false, false, 0, false, false, false},
		'=': {keybd_event.VK_0, true, false, false, 0, false, false, false}, '?': {keybd_event.VK_SP2, true, false, false, 0, false, false, false},
		'¿': {keybd_event.VK_SP3, true, false, false, 0, false, false, false}, '*': {keybd_event.VK_SP5, true, false, false, 0, false, false, false}, 
		';': {keybd_event.VK_SP9, true, false, false, 0, false, false, false}, '+': {keybd_event.VK_SP5, false, false, false, 0, false, false, false},
		':': {keybd_event.VK_SP10, true, false, false, 0, false, false, false}, '_': {keybd_event.VK_SP11, true, false, false, 0, false, false, false},
		'-': {keybd_event.VK_SP11, false, false, false, 0, false, false, false}, '.': {keybd_event.VK_DOT, false, false, false, 0, false, false, false}, 
		',': {keybd_event.VK_COMMA, false, false, false, 0, false, false, false}, '\'': {keybd_event.VK_SP2, false, false, false, 0, false, false, false},
		'¡': {keybd_event.VK_SP3, false, false, false, 0, false, false, false}, 'ª': {keybd_event.VK_SP1, true, false, false, 0, false, false, false},
		'<': {keybd_event.VK_SP12, false, false, false, 0, false, false, false}, '>': {keybd_event.VK_SP12, true, false, false, 0, false, false, false},
		'º': {keybd_event.VK_SP1, false, false, false, 0, false, false, false},
		// Necesitan un espacio en blanco para funcionar
		'^': {keybd_event.VK_SP4, true, false, true, 0, false, false, false}, '¨': {keybd_event.VK_SP7, true, false, true, 0, false, false, false},
		'`': {keybd_event.VK_SP4, false, false, true, 0, false, false, false}, '´': {keybd_event.VK_SP7, false, false, true, 0, false, false, false},
		// Necesitan un acento para funcionar
		'á': {keybd_event.VK_A, false, false, false, 1, false, false, false}, 'Á': {keybd_event.VK_A, false, false, false, 1, false, false, true},
		'é': {keybd_event.VK_E, false, false, false, 1, false, false, false}, 'É': {keybd_event.VK_E, false, false, false, 1, false, false, true},
		'í': {keybd_event.VK_I, false, false, false, 1, false, false, false}, 'Í': {keybd_event.VK_I, false, false, false, 1, false, false, true},
		'ó': {keybd_event.VK_O, false, false, false, 1, false, false, false}, 'Ó': {keybd_event.VK_O, false, false, false, 1, false, false, true},
		'ú': {keybd_event.VK_U, false, false, false, 1, false, false, false}, 'Ú': {keybd_event.VK_U, false, false, false, 1, false, false, true},
        'à': {keybd_event.VK_A, false, false, false, 2, false, false, false}, 'À': {keybd_event.VK_A, false, false, false, 2, false, false, true},
		'è': {keybd_event.VK_E, false, false, false, 2, false, false, false}, 'È': {keybd_event.VK_E, false, false, false, 2, false, false, true},
		'ì': {keybd_event.VK_I, false, false, false, 2, false, false, false}, 'Ì': {keybd_event.VK_I, false, false, false, 2, false, false, true},
		'ò': {keybd_event.VK_O, false, false, false, 2, false, false, false}, 'Ò': {keybd_event.VK_O, false, false, false, 2, false, false, true},
		'ù': {keybd_event.VK_U, false, false, false, 2, false, false, false}, 'Ù': {keybd_event.VK_U, false, false, false, 2, false, false, true},
		// Necesitan altGr para funcionar
		'\\': {keybd_event.VK_SP1, false, true, false, 0, false, false, false}, '|': {keybd_event.VK_1, false, true, false, 0, false, false, false},
		'@': {keybd_event.VK_2, false, true, false, 0, false, false, false}, '#': {keybd_event.VK_3, false, true, false, 0, false, false, false},
        '€': {keybd_event.VK_5, false, true, false, 0, false, false, false}, '¬': {keybd_event.VK_6, false, true, false, 0, false, false, false},
		'~': {keybd_event.VK_4, false, true, true, 0, false, false, false},
		'[': {keybd_event.VK_SP4, false, true, false, 0, false, false, false}, ']': {keybd_event.VK_SP5, false, true, false, 0, false, false, false},
		'{': {keybd_event.VK_SP7, false, true, false, 0, false, false, false}, '}': {keybd_event.VK_SP8, false, true, false, 0, false, false, false},
		// Necesitan caret para funcionar
		'â': {keybd_event.VK_A, false, false, false, 0, true, false, false}, 'Â': {keybd_event.VK_A, false, false, false, 0, true, false, true},
		'ê': {keybd_event.VK_E, false, false, false, 0, true, false, false}, 'Ê': {keybd_event.VK_E, false, false, false, 0, true, false, true},
		'î': {keybd_event.VK_I, false, false, false, 0, true, false, false}, 'Î': {keybd_event.VK_I, false, false, false, 0, true, false, true},
		'ô': {keybd_event.VK_O, false, false, false, 0, true, false, false}, 'Ô': {keybd_event.VK_O, false, false, false, 0, true, false, true},
		'û': {keybd_event.VK_U, false, false, false, 0, true, false, false}, 'Û': {keybd_event.VK_U, false, false, false, 0, true, false, true},
		// Necesitan dieresis para funcionar
		'ä': {keybd_event.VK_A, false, false, false, 0, false, true, false}, 'Ä': {keybd_event.VK_A, false, false, false, 0, false, true, true},
		'ë': {keybd_event.VK_E, false, false, false, 0, false, true, false}, 'Ë': {keybd_event.VK_E, false, false, false, 0, false, true, true},
		'ï': {keybd_event.VK_I, false, false, false, 0, false, true, false}, 'Ï': {keybd_event.VK_I, false, false, false, 0, false, true, true},
		'ö': {keybd_event.VK_O, false, false, false, 0, false, true, false}, 'Ö': {keybd_event.VK_O, false, false, false, 0, false, true, true},
		'ü': {keybd_event.VK_U, false, false, false, 0, false, true, false}, 'Ü': {keybd_event.VK_U, false, false, false, 0, false, true, true},
	}

	if val, found := mapeo[char]; found {
		return val.key, val.shiftRequired, val.altGRRequired, val.spaceRequired, val.accentoRequired, val.caretRequired, val.dieresisRequired, val.mayuscula
	}
	return keybd_event.VK_3, true , false , false, 0, false, false, false // Caracter no soportado
}