package main

import "fmt"

func main() {
	if nombre, edad := "Pachi", 26; nombre == "Pachi" {
		fmt.Println("Hola", nombre)
	} else {
		fmt.Printf("nombre %s - Edad %d \n", nombre, edad)
	}

	// Obtener valor de mapa

	mapa := make(map[string]string)

	mapa["Pachi"] = "alfonso@gmail.com"
	mapa["Alfonso"] = "pachi@gmail.com"

	// Verificar si el valor existe en el mapa
	if correo, ok := mapa["Paachi"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Print("No existe verga")
	}

}
