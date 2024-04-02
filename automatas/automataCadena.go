package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoCadena int

const (
	EstadoInicialCadena EstadoCadena = iota
	EstadoIntermedioCadena
	EstadoFinalCadena
	EstadoErrorCadena
	EstadoNoAceptado
)

type AutomataCadena struct {
	estadoActual EstadoCadena
}

func (a *AutomataCadena) procesarPalabra(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialCadena:
		if simbolo == '?' {
			a.estadoActual = EstadoIntermedioCadena
		} else {
			a.estadoActual = EstadoNoAceptado
		}
	case EstadoIntermedioCadena:
		if simbolo != '?' {
			//sigue en el mismo estado
		} else if simbolo == '?' {
			a.estadoActual = EstadoFinalCadena
		}
	case EstadoFinalCadena:
		a.estadoActual = EstadoErrorCadena
	}
}

func EvaluarCadena(cadena string) (bool, int) {
	automata := &AutomataCadena{}
	var iterator int

	for i, simbolo := range cadena {
		automata.procesarPalabra(simbolo)

		if automata.estadoActual == EstadoNoAceptado {
			break
		}
		// Agregar tu condición de parada aquí
		if automata.estadoActual == EstadoFinalCadena {
			iterator = i + 1
			break
		}
		if i == len(cadena)-1 && simbolo != '?'{
			automata.estadoActual = EstadoErrorCadena
			iterator = i +1
			break
		}
	}



	if automata.estadoActual == EstadoFinalCadena {
		contenido := cadena[:iterator] + " -> cadena"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un cadena: ", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorCadena {
		contenido := cadena[:iterator] + " -> error sintactivo cadena"
		fmt.Println("error sintactico cadena-> : ", cadena[:iterator])
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		return true, iterator
	}

	return false, len(cadena)
}
