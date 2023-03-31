package main

import "fmt"

func sumar(numeros ...int) int {
	var num int
	for i := 0; i < len(numeros); i++ {
		num += numeros[i]
	}
	fmt.Println(numeros)
	return num
}

func sumar2(nombre string, numeros ...int) (string, int) {
	var values int

	for _, numeros := range numeros {
		values += numeros
	}

	mensaje := fmt.Sprintf("La suma de %s es de %d", nombre, values)
	return mensaje, values
}

func main() {

	resultado := sumar(4, 5, 8, 6)
	fmt.Println(resultado)

	var myname string

	fmt.Println("Dime tu nombre")

	fmt.Scanln(&myname)

	msg, result := sumar2(myname, 10, 20, 30)
	fmt.Println(msg)
	fmt.Println(result)
}
