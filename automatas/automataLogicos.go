package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoLogico int

const (
	EstadoIncialLogico EstadoLogico = iota
	EstadoFinalLogico
	EstadoErrorLogico
	EstadoNoAceptadoLogico
)

type AutomataLogico struct {
	estadoActual EstadoLogico
}

func (a *AutomataLogico) procesarLogico(simbolo rune) {
	switch a.estadoActual {
	case EstadoIncialLogico:
		if simbolo == '&' {
			a.estadoActual = EstadoFinalLogico
		} else if simbolo == '|' {
			a.estadoActual = EstadoFinalLogico
		} else if simbolo == '¬' {
			a.estadoActual = EstadoFinalLogico
		} else {
			a.estadoActual = EstadoNoAceptadoLogico
		}
	case EstadoFinalLogico:
		a.estadoActual = EstadoErrorLogico
	}
}

func EvaluarLogico(cadena string) (bool, int) {
	automata := &AutomataLogico{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarLogico(simbolo)

		if automata.estadoActual == EstadoNoAceptadoLogico {
			break
		}

		if automata.estadoActual == EstadoFinalLogico {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorLogico {
			iterator = iterator + 1
			break
		}
	}

	if automata.estadoActual == EstadoFinalLogico {
		contenido := cadena[:iterator] + " -> logico"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un operador logico", cadena[:iterator])
		return true, iterator
	}
	return false, len(cadena)

}
