package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

// Definición del tipo de estado
type EstadoS int

const (
	EstadoInicialS EstadoS = iota
	EstadoSegundaS
	EstadoIntermedioS
	EstadoPreFinalS
	EstadoFinalS
	EstadoErrorS
	EstadoNoAceptadoIdentificador
)

type AutomataS struct {
	estadoActual EstadoS
	contador     int
}

func (a *AutomataS) ProcesarSimbolo(simbolo rune) {

	switch a.estadoActual {
	case EstadoInicialS:
		if simbolo == '$' {
			a.estadoActual = EstadoSegundaS
		} else {
			a.estadoActual = EstadoNoAceptadoIdentificador
		}
	case EstadoSegundaS:
		if simbolo == '$' {
			a.estadoActual = EstadoIntermedioS
			a.contador = 1
		} else {
			a.estadoActual = EstadoErrorS
		}
	case EstadoIntermedioS:
		if esCaracterValido(simbolo) {
			if a.contador <= 10 {
				// Permanecer en el estado intermedio mientras no se superen los 10 caracteres
				a.contador++
			} else {
				a.estadoActual = EstadoErrorS
			}
		} else if simbolo == '$' && a.contador <= 10 {
			a.estadoActual = EstadoPreFinalS
		}
	case EstadoPreFinalS:
		if simbolo == '$' {
			a.estadoActual = EstadoFinalS
		} else {
			a.estadoActual = EstadoErrorS
		}
	case EstadoFinalS:
		a.estadoActual = EstadoErrorS
	}
}

func esCaracterValido(simbolo rune) bool {
	// Puedes ajustar la lógica según tus necesidades
	return (simbolo >= 'a' && simbolo <= 'z') || (simbolo >= 'A' && simbolo <= 'Z') || simbolo == ' '
}

func EvaluarIdentificador(cadena string) (bool, int) {
	automata := &AutomataS{}
	var iterator int
	for i, simbolo := range cadena {
		automata.ProcesarSimbolo(simbolo)

		if automata.estadoActual == EstadoNoAceptadoIdentificador {
			break
		}

		if automata.estadoActual == EstadoFinalS {
			iterator = i + 1
			break
		}
		if automata.estadoActual == EstadoErrorS {
			iterator = i + 1
			break
		}

		if i == len(cadena)-1 && simbolo != '$' {
			automata.estadoActual = EstadoErrorS
			iterator = i+1
			break
		}
	}

	if automata.estadoActual == EstadoFinalS {
		contenido := cadena[:iterator] + " -> identificador"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un Identificador", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorS {
		contenido := cadena[:iterator] + " -> error sintactico identificador"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		return true, iterator
	}

	return false, len(cadena)

}
