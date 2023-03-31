package main

import "fmt"

func main() {
	a := 10
	b := 20

	r := a == b
	fmt.Printf("%d es igual a %d ? : %t", a, b, r)

	r = a != b
	fmt.Printf("\n%d es diferente a %d ? : %t", a, b, r)

	r = a > b
	fmt.Printf("\n%d es mayor que %d ? : %t", a, b, r)

	r = a < b
	fmt.Printf("\n%d es menor que %d ? : %t", a, b, r)

	r = a >= b
	fmt.Printf("\n%d es mayor o igual que %d ? : %t", a, b, r)

	r = a <= b
	fmt.Printf("\n%d es menor o igual que %d ? : %t", a, b, r)
}
