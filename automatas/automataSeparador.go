package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoSeparador int

const (
	EstadoInicialSeparador EstadoSeparador = iota
	EstadoFinalSeparador
	EstadoNoAceptadoSeparador
)

type AutomataSeparador struct {
	estadoActual EstadoSeparador
}

func (a *AutomataSeparador) procesarSeparador(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialSeparador:
		if simbolo == ',' {
			a.estadoActual = EstadoFinalSeparador
		} else {
			a.estadoActual = EstadoNoAceptadoSeparador
		}
	case EstadoFinalSeparador:
		a.estadoActual = EstadoNoAceptadoSeparador
	}
}

func EvaluarSeparador(cadena string) (bool, int) {
	automata := &AutomataSeparador{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarSeparador(simbolo)
		if automata.estadoActual == EstadoNoAceptadoSeparador {
			break
		}

		if automata.estadoActual == EstadoFinalSeparador {
			iterator = i + 1
			break
		}
	}

	if automata.estadoActual == EstadoFinalSeparador {
		contenido := cadena[:iterator] + " -> separador"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un separador: ", cadena[:iterator])
		return true, iterator
	}
	return false, len(cadena)

}
