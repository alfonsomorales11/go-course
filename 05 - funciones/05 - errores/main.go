package main

import (
	"errors"
	"fmt"
)

func division(dividendo int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre cero")
	} else {
		return dividendo / divisor, nil
	}
}

func main() {
	resultado, error := division(10, 0)

	if error == nil {
		fmt.Println("El resultado es: ", resultado)
	} else {
		fmt.Println(error)
	}
}
