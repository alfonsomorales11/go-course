package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if error := recover(); error != nil {
			panic("Ocurrio un error inesperado")
		}
	}()

	if file, error := os.Open("hola.txt"); error != nil {
		panic("Ocurrio un error")
	} else {

		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

}
