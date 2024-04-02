package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoReales int

const (
	EstadoInicial EstadoReales = iota
	EstadoR
	EstadoNumerosReal
	EstadoComaReal
	EstadoDespuesComaReal
	EstadoFinalReal
	EstadoErrorReal
	EstadoNoAceptadoReal
)

type AutomataNumeroReal struct {
	estadoActual EstadoReales
}

func (a *AutomataNumeroReal) ProcesarSimbolo(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicial:
		if simbolo == 'r' {
			a.estadoActual = EstadoR
		} else {
			a.estadoActual = EstadoNoAceptadoReal
		}
	case EstadoR:
		if esDigito(simbolo) {
			a.estadoActual = EstadoNumerosReal
		} else {
			a.estadoActual = EstadoErrorReal
		}
	case EstadoNumerosReal:
		if esDigito(simbolo) {
			// Permanecer en el estado de números
		} else if simbolo == ',' {
			a.estadoActual = EstadoComaReal
		} else {
			a.estadoActual = EstadoErrorReal
		}
	case EstadoComaReal:
		if simbolo == '0' || esDigito(simbolo) {
			a.estadoActual = EstadoDespuesComaReal
		} else {
			a.estadoActual = EstadoErrorReal
		}
	case EstadoDespuesComaReal:
		if esDigito(simbolo) {
			// Permanecer en el estado después de la coma
		} else if simbolo == 'r' {
			a.estadoActual = EstadoFinalReal
		} else {
			a.estadoActual = EstadoErrorReal
		}
	case EstadoFinalReal:
		a.estadoActual = EstadoErrorReal
	}
}

func esDigito(simbolo rune) bool {
	return simbolo >= '0' && simbolo <= '9'
}

func EvaluarReales(cadena string) (bool, int) {
	automata := &AutomataNumeroReal{}
	var iterator int
	for i, simbolo := range cadena {
		automata.ProcesarSimbolo(simbolo)

		if automata.estadoActual == EstadoNoAceptadoReal {
			break
		}

		if automata.estadoActual == EstadoFinalReal {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorReal {
			iterator = i + 1
			break
		}

		if i == len(cadena)-1 && simbolo != 'r' {
			automata.estadoActual = EstadoErrorReal
			iterator = i
			break
		}
	}

	if automata.estadoActual == EstadoFinalReal {
		contenido := cadena[:iterator] + " -> numero real"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un numero real ", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorReal {
		contenido := cadena[:iterator] + " -> error sintactico numero real"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("error sintactico numero real ", cadena[:iterator])
		return true, iterator
	}

	return false, len(cadena)

}
