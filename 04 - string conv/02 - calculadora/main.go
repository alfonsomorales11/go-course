package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	fmt.Println(valores)

	var result int

	for _, valor := range valores {
		num, error := strconv.Atoi(valor)

		if error != nil {
			fmt.Println(error)
			fmt.Println("Tiene que ingresar un numero entero")
			fmt.Println("Solo debe ingresar un operador")
		} else {
			result += num
		}
	}

	return result
}

func main() {
	var expresion string
	// var result int

	fmt.Println("=")
	fmt.Scanln(&expresion)

	result := sumar(expresion)
	fmt.Println(result)

}
