package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoIncrementos int

const (
	EstadoInicialIncremento EstadoIncrementos = iota
	EstadoInicialSumIncr
	EstadoInicialResIncr
	EstadoFinalIncremento
	EstadoErrorIncremento
	EstadoNoAceptadoIncremento
)

type AutomataIncremento struct {
	estadoActual EstadoIncrementos
}

func (a *AutomataIncremento) procesarIncremento(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialIncremento:
		if simbolo == '{' {
			a.estadoActual = EstadoInicialSumIncr
		} else if simbolo == '}' {
			a.estadoActual = EstadoInicialResIncr
		} else {
			a.estadoActual = EstadoNoAceptadoIncremento
		}
	case EstadoInicialResIncr:
		if simbolo == '}' {
			a.estadoActual = EstadoFinalIncremento
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoInicialSumIncr:
		if simbolo == '{' {
			a.estadoActual = EstadoFinalIncremento
		} else {
			a.estadoActual = EstadoErrorIncremento
		}
	case EstadoFinalIncremento:
		a.estadoActual = EstadoErrorIncremento
	}
}

func EvaluarIncrementos(cadena string) (bool, int) {
	automata := &AutomataIncremento{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarIncremento(simbolo)

		if automata.estadoActual == EstadoNoAceptadoIncremento {
			break
		}

		if automata.estadoActual == EstadoFinalIncremento {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorIncremento {
			iterator = i + 1
			break
		}
	}

	if automata.estadoActual == EstadoFinalIncremento {
		contenido := cadena[:iterator] + " -> incremento"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un operador de incremento", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorIncremento {
		contenido := cadena[:iterator] + " -> error sintactico incremetno"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("error sintactico incremento ", cadena[:iterator])
		return true, iterator
	}

	return false, len(cadena)

}
