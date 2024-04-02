package main

import (
	"fmt"
	"proyecto_tlf/automatas"
	"proyecto_tlf/utilities"
)

func main() {
	utilities.EliminarArchivo("./salida.txt")
	texto := utilities.LeerArchivo("./entrada.txt")
	var txtOriginal string
	var recorrido int
	recorrido = 0
	txtOriginal = texto
	fin := len(txtOriginal)

	for i := 0; recorrido < fin; i++ {
		if recorrido < fin {
			texto = txtOriginal[recorrido:fin]
			posicion := esCadena(texto)

			if posicion != -1 {

				recorrido += posicion
			} else {
				posicion = esAsignacion(texto)
				if posicion != -1 {
					recorrido += posicion
				} else {
					posicion = esAritmetico(texto)

					if posicion != -1 {

						recorrido += posicion
					} else {
						posicion = esLogico(texto)
						if posicion != -1 {
							recorrido += posicion
						} else {
							posicion = esNatural(texto)
							if posicion != -1 {

								recorrido += posicion
							} else {

								posicion = esReal(texto)
								if posicion != -1 {

									recorrido += posicion
								} else {

									posicion = esComparacion(texto)
									if posicion != -1 {

										recorrido += posicion
									} else {
										posicion = esIncremento(texto)
										if posicion != -1 {

											recorrido += posicion
										} else {
											posicion = esTerminal(texto)
											if posicion != -1 {

												recorrido += posicion
											} else {
												posicion = esSeparador(texto)
												if posicion != -1 {
													fmt.Println("entro aqui")
													recorrido += posicion
												} else {
													posicion = esHexadecimal(texto)
													if posicion != -1 {
														recorrido += posicion
													} else {
														posicion = esComentario(texto)
														if posicion != -1 {
															recorrido += posicion
														} else {
															posicion = esIdentificador(texto)
															if posicion != -1 {
																recorrido += posicion
															} else {
																posicion = esPalabraReservada(texto)
																if posicion != -1 {
																	recorrido += posicion
																} else {
																	// error
																	if txtOriginal[recorrido] == ' ' {
																		recorrido++
																	} else {
																		recorrido++
																		contenido := txtOriginal[recorrido-1:recorrido] + " -> token no reconocido"
																		utilities.GuardarEnArchivo(contenido, "./salida.txt")

																	}
																}

															}

														}
													}
												}
											}
										}
									}

								}

							}

						}

					}
				}
			}

		}
	}

}

func esPalabraReservada(texto string) int {
	valido, tam := automatas.EvaluarPR(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esCadena(texto string) int {
	valido, tam := automatas.EvaluarCadena(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esAsignacion(texto string) int {
	valido, tam := automatas.EvaluarIgualacion(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esAritmetico(texto string) int {
	valido, tam := automatas.EvaluarAritmetico(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esLogico(texto string) int {
	valido, tam := automatas.EvaluarLogico(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esNatural(texto string) int {
	valido, tam := automatas.EvaluarNatural(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esReal(texto string) int {
	valido, tam := automatas.EvaluarReales(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esComparacion(texto string) int {
	valido, tam := automatas.EvaluarComparacion(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esIncremento(texto string) int {
	valido, tam := automatas.EvaluarIncrementos(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esTerminal(texto string) int {
	valido, tam := automatas.EvaluarTerminal(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esSeparador(texto string) int {
	valido, tam := automatas.EvaluarSeparador(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esHexadecimal(texto string) int {
	valido, tam := automatas.EvalaurHexadecimal(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esComentario(texto string) int {
	valido, tam := automatas.EvaluarComentario(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}

func esIdentificador(texto string) int {
	valido, tam := automatas.EvaluarIdentificador(texto)
	if valido {
		return tam
	} else {
		return -1
	}
}
