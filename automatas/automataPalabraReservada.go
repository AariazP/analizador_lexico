package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type estadoPalabraReservada int

const (
	EstadoInicialPalabraReservada estadoPalabraReservada = iota

	EstadoInicioPalabra

	EstadoFinM
	EstadoFinParCIe
	EstadoFinD
	EstadoFinParArr

	EstadoFinalPr

	EstadoNoSoportadoPR
	EstadoErrorPr
)

type AutomataPR struct {
	estadoActual estadoPalabraReservada
}

func (a *AutomataPR) procesarPalabraReservada(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialPalabraReservada:
		if simbolo == 'U' || simbolo == 'Y' {
			a.estadoActual = EstadoFinM
		} else if simbolo == 'T' {
			a.estadoActual = EstadoFinParCIe
		} else if simbolo == 'F' || simbolo == 'G' {
			a.estadoActual = EstadoFinD
		} else if simbolo == 'V' {
			a.estadoActual = EstadoFinParArr
		} else {
			a.estadoActual = EstadoNoSoportadoPR
		}
	case EstadoFinM:
		if simbolo == 'M' {
			a.estadoActual = EstadoFinalPr
		} else {
			a.estadoActual = EstadoErrorPr
		}
	case EstadoFinParCIe:
		if simbolo == ')' {
			a.estadoActual = EstadoFinalPr
		} else {
			a.estadoActual = EstadoErrorPr
		}
	case EstadoFinD:
		if simbolo == 'D' {
			a.estadoActual = EstadoFinalPr
		} else {
			a.estadoActual = EstadoErrorPr
		}
	case EstadoFinParArr:
		if simbolo == '(' {
			a.estadoActual = EstadoFinalPr
		} else {
			a.estadoActual = EstadoErrorPr
		}
	}
}

func EvaluarPR(cadena string) (bool, int) {
	automata := &AutomataPR{}
	var iterator int

	for i, simbolo := range cadena {
		automata.procesarPalabraReservada(simbolo)

		if automata.estadoActual == EstadoNoSoportadoPR {
			break
		}

		if automata.estadoActual == EstadoFinalPr {
			iterator = i + 1
			break
		}

		if automata.estadoActual == EstadoErrorPr {
			iterator = i + 1
			break
		}

		x := simbolo != 'M'
		x1 := simbolo != 'H'
		x2 := simbolo != 'D'
		if i == len(cadena)-1 && (x || x1 || x2 || simbolo != 'B') {
			automata.estadoActual = EstadoErrorPr
			iterator = i
			break
		}
	}

	if automata.estadoActual == EstadoFinalPr {
		contenido := cadena[:iterator] + " -> palabra reservada"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorPr {
		contenido := cadena[:iterator] + " -> error sintactico palabra reservada"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("error sintactico palabra reservada -> ", cadena[:iterator])
		return true, iterator
	}
	return false, len(cadena)
}
