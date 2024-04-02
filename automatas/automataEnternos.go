package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

// Definición del tipo de estado
type EstadoNatural int

const (
	EstadoInicialNatural EstadoNatural = iota
	EstadoNumeroNatural
	EstadoFinalNatural
	EstadoErrorNatural
	EstadoNoAceptadoNatural
)

type AutomataNatural struct {
	estadoActual EstadoNatural
}

func (a *AutomataNatural) ProcesarSimbolo(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialNatural:
		if simbolo == 'n' {
			a.estadoActual = EstadoNumeroNatural
		} else {
			a.estadoActual = EstadoNoAceptadoNatural
		}
	case EstadoNumeroNatural:
		if esDigito(simbolo) {
			// Permanecer en el estado de número
		} else if simbolo == 'n' {
			a.estadoActual = EstadoFinalNatural
		} else {
			a.estadoActual = EstadoErrorNatural
		}
	case EstadoFinalNatural:
		a.estadoActual = EstadoErrorNatural
	}
}

func esDigitoNatural(simbolo rune) bool {
	return simbolo >= '0' && simbolo <= '9'
}

func EvaluarNatural(cadena string) (bool, int) {
	automata := &AutomataNatural{}
	var iterator int
	for i, simbolo := range cadena {
		automata.ProcesarSimbolo(simbolo)

		if automata.estadoActual == EstadoNoAceptadoNatural {
			break
		}

		if automata.estadoActual == EstadoFinalNatural {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorNatural {
			//fmt.Println("entro aqui")
			iterator = i + 1
			break
		}

		if i == len(cadena)-1 && simbolo != 'n' {
			automata.estadoActual = EstadoErrorNatural
			iterator = i 
			break
		}
	}

	if automata.estadoActual == EstadoFinalNatural {
		contenido := cadena[:iterator] + " -> numero natural"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un numero natural", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorNatural {
		contenido := cadena[:iterator] + " -> error sintactico natural"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("error sintactico numero natural -> ", cadena[:iterator])
		return true, iterator
	}

	return false, len(cadena)

}
