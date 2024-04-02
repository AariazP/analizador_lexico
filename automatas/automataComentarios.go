package automatas

import (
	"fmt"
	"proyecto_tlf/utilities"
)

type EstadoComentario int

const (
	EstadoInicialComentarios EstadoComentario = iota
	//comentario de linea
	EstadoInicioComLinea
	EstadoIntermedioComLinea

	//comentario de bloque
	EstadoInicioComBlo
	EstadoComBloIntermedio
	EstadoComBloPosIntermedio

	//estado final
	EstadoFinalComentario
	EstadoErrorComentario

	EstadoNoAceptadoComentario
)

type AutomataComentario struct {
	estadoActual EstadoComentario
}

func (a *AutomataComentario) procesarComentario(simbolo rune) {
	switch a.estadoActual {
	case EstadoInicialComentarios:
		if simbolo == '#' {
			a.estadoActual = EstadoInicioComLinea
		} else if simbolo == '/' {
			a.estadoActual = EstadoInicioComBlo
		} else {
			a.estadoActual = EstadoNoAceptadoComentario
		}
	case EstadoInicioComLinea:
		terminal, _ := EvaluarTerminal(string(simbolo))
		if terminal == true {
			a.estadoActual = EstadoFinalComentario
		} else {
			a.estadoActual = EstadoInicioComLinea
		}
	case EstadoInicioComBlo:
		if simbolo == '#' {
			a.estadoActual = EstadoComBloIntermedio
		} else {
			a.estadoActual = EstadoErrorComentario
		}
	case EstadoComBloIntermedio:
		if simbolo != '#' {
			//sigue en el mismo estado
		} else if simbolo == '#' {
			a.estadoActual = EstadoComBloPosIntermedio
		} else {
			a.estadoActual = EstadoErrorComentario
		}
	case EstadoComBloPosIntermedio:
		if simbolo == '/' {
			a.estadoActual = EstadoFinalComentario
		} else {
			a.estadoActual = EstadoErrorComentario
		}
	case EstadoFinalComentario:
		a.estadoActual = EstadoErrorComentario
	}

}

func EvaluarComentario(cadena string) (bool, int) {
	automata := &AutomataComentario{}
	var iterator int
	for i, simbolo := range cadena {
		automata.procesarComentario(simbolo)

		if automata.estadoActual == EstadoNoAceptadoComentario {
			break
		}

		if automata.estadoActual == EstadoFinalComentario {

			iterator = i + 1
			break
		}

		if automata.estadoActual == EstadoErrorComentario {
			iterator = i + 1
			break
		}

		x := simbolo != ';'
		if i == len(cadena)-1 && (x || simbolo != '/') {
			automata.estadoActual = EstadoErrorComentario
			iterator = i + 1
			break
		}
	}

	
	if automata.estadoActual == EstadoFinalComentario {
		contenido := cadena[:iterator] + " -> comentario"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		fmt.Println("es un comentario", cadena[:iterator])
		return true, iterator
	}

	if automata.estadoActual == EstadoErrorComentario {

		contenido := cadena[:iterator] + " -> erros sintactico comentario"
		utilities.GuardarEnArchivo(contenido, "./salida.txt")
		return true, iterator
	}
	return false, len(cadena)

}
