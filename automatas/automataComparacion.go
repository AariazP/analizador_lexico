package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoComparacion int

const (
	EstadoInicialComparacion EstadoComparacion = iota
	EstadoIntermedioComparacion
	EstadoIntermedioIgualacion
	EstadoPosIntermedioComparacion
	EstadoFinalComparacion
	EstadoErrorComparacion
	EstadoNoAceptadoComparacion
)

type AutomataComparacion struct {
	estadoActual EstadoComparacion
}

func (a *AutomataComparacion) procesarPalabra(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialComparacion:
		if simbolo == 'M' {
			a.estadoActual = EstadoIntermedioComparacion
		} else if simbolo == 'E' {
			a.estadoActual = EstadoIntermedioIgualacion
		} else {
			a.estadoActual = EstadoNoAceptadoComparacion
		}
	case EstadoIntermedioComparacion:
		if simbolo == 'Q' || simbolo == 'N' {
			a.estadoActual = EstadoFinalComparacion
		} else {
			a.estadoActual = EstadoErrorComparacion
		}
	case EstadoPosIntermedioComparacion:
		if simbolo == ' ' {
			a.estadoActual = EstadoFinalComparacion
		} else if simbolo == 'E' {
			a.estadoActual = EstadoIntermedioIgualacion
		}
	case EstadoIntermedioIgualacion:
		if simbolo == 'Q' {
			a.estadoActual = EstadoFinalComparacion
		} else {
			a.estadoActual = EstadoErrorComparacion
		}
	case EstadoFinalComparacion:
		if simbolo == 'E' {
			a.estadoActual = EstadoIntermedioIgualacion
		} else {
			a.estadoActual = EstadoErrorComparacion
		}
	}

}

func EvaluarComparacion(cadena string) (bool, int) {
	automata := &AutomataComparacion{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarPalabra(simbolo)

		if automata.estadoActual == EstadoNoAceptadoComparacion {
			break
		}

		if automata.estadoActual == EstadoFinalComparacion {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorComparacion {
			iterator = i + 1
			break
		}
	}

	if automata.estadoActual == EstadoFinalComparacion {
		contenido := cadena[:iterator] + " -> comparacion"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un operador de comparacion ", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorComparacion {

		contenido := cadena[:iterator] + " -> error sintactico comparacion"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("error sintactico comparacion")
		return true, iterator

	}

	return false, len(cadena)

}
