package utilities

import (
	"io/ioutil"
)

func LeerArchivo(nombreArchivo string) string {
	// Lee el contenido del archivo
	contenido, _ := ioutil.ReadFile(nombreArchivo)

	// Convierte el contenido a un string y lo retorna
	return string(contenido)
}

