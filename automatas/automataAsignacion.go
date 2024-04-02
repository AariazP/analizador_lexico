package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoIgualacion int

const (
	EstadoIncialIgualacion EstadoIgualacion = iota
	EstadoIntermedioIgual
	EstadoFinalIgualacion
	EstadoErrorIgualacion
	EstadoNoAceptadoIgualacion
)

type AutomataIgualacion struct {
	estadoActual EstadoIgualacion
}

func (a *AutomataIgualacion) procesarIgualacion(simbolo rune) {
	switch a.estadoActual {
	case EstadoIncialIgualacion:
		if simbolo == '.' {
			a.estadoActual = EstadoIntermedioIgual
		} else {
			a.estadoActual = EstadoNoAceptadoIgualacion
		}
	case EstadoIntermedioIgual:
		if simbolo != '=' {
			a.estadoActual = EstadoErrorIgualacion
		} else {
			a.estadoActual = EstadoFinalIgualacion
		}
	case EstadoFinalIgualacion:
		a.estadoActual = EstadoErrorIgualacion
	}
}

func EvaluarIgualacion(cadena string) (bool, int) {
	automata := &AutomataIgualacion{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarIgualacion(simbolo)

		if automata.estadoActual == EstadoNoAceptadoIgualacion{
			break
		}

		if automata.estadoActual == EstadoFinalIgualacion || automata.estadoActual == EstadoErrorIgualacion {
			iterator = i + 1
			break
		}

		if i == len(cadena)-1 && simbolo != '=' {
			automata.estadoActual = EstadoErrorIgualacion
			iterator = i 
			break
		}

	}

	if automata.estadoActual == EstadoFinalIgualacion {
		contenido := cadena[:iterator] + " -> asignacion"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un asignacion: ", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorIgualacion {
		contenido := cadena[:iterator] + " -> error sintactico asignacion"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("error sintactico asignacion -> ", cadena[:iterator])
		return true, iterator
	}

	return false, len(cadena)
}
