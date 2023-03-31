package main

import "fmt"

func main() {

	var numeros [5]int

	numeros[0] = 10
	numeros[1] = 20

	fmt.Println(numeros)

	nombres := [3]string{"Alfonso", "Morales", "Pachimix"}

	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Verde",
		"Azul",
	}

	fmt.Println(colores, len(colores))

	monedas := [...]string{
		0: "dolar",
		2: "soles",
		3: "pesos",
		5: "euros",
	}

	fmt.Println(monedas, len(monedas))

	subMoneda := monedas[0:3]
	fmt.Println(subMoneda)
}
