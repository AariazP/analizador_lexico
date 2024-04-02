package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoTermial int

const (
	EstadoInicialTerminal EstadoTermial = iota
	EstadoIntermedioTerminal
	EstadoFinalTermianl
	EstadoErrorTerminal
	EstadoNoAceptadoTerminal
)

type AutomataTerminal struct {
	estadoActual EstadoTermial
}

func (a *AutomataTerminal) procesarTerminal(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialTerminal:
		if simbolo == ';' {
			a.estadoActual = EstadoFinalTermianl
		} else if simbolo == '\n' {
			a.estadoActual = EstadoFinalTermianl
		} else {
			a.estadoActual = EstadoNoAceptadoTerminal
		}
	case EstadoFinalTermianl:
		a.estadoActual = EstadoErrorTerminal
	}
}

func EvaluarTerminal(cadena string) (bool, int) {
	automata := &AutomataTerminal{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarTerminal(simbolo)

		if automata.estadoActual == EstadoNoAceptadoTerminal {
			break
		}

		if automata.estadoActual == EstadoFinalTermianl {
			iterator = i + 1
			break
		}
	}

	if automata.estadoActual == EstadoFinalTermianl {
		contenido := cadena[:iterator] + " -> terminal"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un operador terminal", cadena[:iterator])
		return true, iterator
	}
	return false, len(cadena)

}
