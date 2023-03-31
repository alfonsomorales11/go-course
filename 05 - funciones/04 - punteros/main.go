package main

import "fmt"

//						* indica que es una referencia
func modificarValor(cadena *string) {
	fmt.Printf("%p \n", &cadena)
	// *variableName modifica el valor de un puntero
	*cadena = "hola desde la funcion"
}

func main() {
	cadena := "hola mundo de go\n"
	fmt.Printf("%p", &cadena)
	fmt.Println("\nAntes de la funcion: ", cadena)

	//             &cadena tiene como valor la direccion en memoria de la variable
	modificarValor(&cadena)
	fmt.Println("Despues de la funcion: ", cadena)
}
