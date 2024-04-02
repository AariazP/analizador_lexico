package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoAritmetico int

const (
	EstadoInicialArimetico EstadoAritmetico = iota
	EstadoIncialSuma
	EstadoIncialResta
	EstadoIncialMultiplicacion
	EstadoIncialDivision
	EstadoFInalAritmetico
	EstadoErrorAritmetico
	EstadoNoAceptadoAritmetico
)

type AutomataArimetico struct {
	estadoActual EstadoAritmetico
}

func (a *AutomataArimetico) procesarArimetico(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialArimetico:
		if simbolo == 'm' {
			a.estadoActual = EstadoIncialSuma
		} else if simbolo == 's' {
			a.estadoActual = EstadoIncialResta
		} else if simbolo == 'p' {
			a.estadoActual = EstadoIncialMultiplicacion
		} else if simbolo == 'd' {
			a.estadoActual = EstadoIncialDivision
		} else {
			a.estadoActual = EstadoNoAceptadoAritmetico
		}
	case EstadoIncialSuma:
		if simbolo == 's' {
			a.estadoActual = EstadoFInalAritmetico
		} else {
			a.estadoActual = EstadoErrorAritmetico
		}
	case EstadoIncialResta:
		if simbolo == 'm' {
			a.estadoActual = EstadoFInalAritmetico
		} else {
			a.estadoActual = EstadoErrorAritmetico
		}
	case EstadoIncialMultiplicacion:
		if simbolo == 'r' {
			a.estadoActual = EstadoFInalAritmetico
		} else {
			a.estadoActual = EstadoErrorAritmetico
		}
	case EstadoIncialDivision:
		if simbolo == 'v' {
			a.estadoActual = EstadoFInalAritmetico
		} else {
			a.estadoActual = EstadoErrorAritmetico
		}

	case EstadoFInalAritmetico:
		a.estadoActual = EstadoErrorAritmetico
	}
}

func EvaluarAritmetico(cadena string) (bool, int) {
	automata := &AutomataArimetico{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarArimetico(simbolo)

		if automata.estadoActual == EstadoNoAceptadoAritmetico{
			break
		}

		if automata.estadoActual == EstadoFInalAritmetico {
			iterator = i + 1
			break
		} else if automata.estadoActual == EstadoErrorAritmetico {
			iterator = i+1
			break
		}
	}

	if automata.estadoActual == EstadoFInalAritmetico {
		contenido := cadena[:iterator] + " -> aritmetico"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un operador aritmetico", cadena[:iterator])
		return true, iterator
	} 

	if automata.estadoActual == EstadoErrorAritmetico {
		contenido := cadena[:iterator] + " -> error sintactico aritmetico"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("error sintactico aritmetico ->", cadena[:iterator])
		return true, iterator
	}
	return false, len(cadena)
	
}
