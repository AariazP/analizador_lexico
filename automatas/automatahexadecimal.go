package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoHexadecimal int

const (
	EstadoInicialHexadecimal EstadoHexadecimal = iota
	EstadoIntermedioHexadecimal
	EstadoFinalHexadecimal
	EstadoNoAceptadoHexadecimal
	EstadoErrorHexadecimal
)

type AutomataHexadecimal struct {
	estadoActual EstadoHexadecimal
}

func (a *AutomataHexadecimal) procesarHexadecimal(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialHexadecimal:
		if simbolo == 'h' {
			a.estadoActual = EstadoIntermedioHexadecimal
		} else {
			a.estadoActual = EstadoNoAceptadoHexadecimal
		}
	case EstadoIntermedioHexadecimal:
		if esMayusculaValida(simbolo) {
			fmt.Println("sigue aqui")
			a.estadoActual = EstadoIntermedioHexadecimal
		} else if simbolo == 'h' {
			fmt.Println("llego al final")
			a.estadoActual = EstadoFinalHexadecimal
		} else {
			a.estadoActual = EstadoIntermedioHexadecimal
		}
	case EstadoFinalHexadecimal:
		a.estadoActual = EstadoErrorHexadecimal

	}
}

func esMayusculaValida(simbolo rune) bool {
	return (simbolo >= 'A' && simbolo <= 'F')
}

func EvalaurHexadecimal(cadena string) (bool, int) {
	automata := &AutomataHexadecimal{}
	var iterator int

	for i, simbolo := range cadena {
		fmt.Println(simbolo)
		automata.procesarHexadecimal(simbolo)

		if automata.estadoActual == EstadoNoAceptadoHexadecimal {
			break
		}

		if automata.estadoActual == EstadoFinalHexadecimal {
			iterator = i + 1
			break
		}

		if i == len(cadena)-1 && i != 'h' {
			automata.estadoActual = EstadoErrorHexadecimal
			iterator = i + 1
			break
		}
	}

	if automata.estadoActual == EstadoFinalHexadecimal {
		contenido := cadena[:iterator] + " -> hexadecimal"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un hexadecimal: ", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorHexadecimal {
		contenido := cadena[:iterator] + " -> error sintactivo hexadecimal"
		fmt.Println("error sintactico -> : ", cadena[:iterator])
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		return true, iterator
	}

	return false, len(cadena)
}
