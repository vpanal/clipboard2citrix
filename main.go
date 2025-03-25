package main

import (
	"fmt"
	"unicode"
	//"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/vpanal/keybd_event"
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

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		if debugging {
			fmt.Println("Error al inicializar el teclado virtual:", err)
		}
		return
	}

	// Simula Alt+Tab para cambiar de ventana
	kb.SetKeys(keybd_event.VK_TAB)
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
func escribirCaracter(kb keybd_event.KeyBonding, char rune) error {
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
		return keybd_event.VK_KP0
	case '1':
		return keybd_event.VK_KP1
	case '2':
		return keybd_event.VK_KP2
	case '3':
		return keybd_event.VK_KP3
	case '4':
		return keybd_event.VK_KP4
	case '5':
		return keybd_event.VK_KP5
	case '6':
		return keybd_event.VK_KP6
	case '7':
		return keybd_event.VK_KP7
	case '8':
		return keybd_event.VK_KP8
	case '9':
		return keybd_event.VK_KP9
	}
	return -1
}
