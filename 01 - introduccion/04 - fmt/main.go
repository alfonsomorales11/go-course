package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Print(hola)
	fmt.Print(mundo)

	nombre := "Alfonso"
	edad := 26

	fmt.Printf("\nHola %s , tu edad es %d años", nombre, edad)
	fmt.Printf("\nHola %v , tu edad es %v años", nombre, edad)

	mensaje := fmt.Sprintf("\nHola %v , tu edad es %v años", nombre, edad)
	fmt.Println(mensaje)

	fmt.Printf("nombre: %T", nombre)

	var nombre2 string
	fmt.Print("\nIngresa otro nombre: ")
	fmt.Scan(&nombre2)

	fmt.Printf("El otro nombre es %s", nombre2)
}
