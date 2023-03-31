package main

import "fmt"

func saludar(nombre string) {
	fmt.Printf("Hola %s\n", nombre)
}

func sumar(a int, b int) int {
	var c int = a + b
	return c
}

func main() {
	saludar("Pachi")

	var result int = sumar(10, 20)
	fmt.Println(result)
}
