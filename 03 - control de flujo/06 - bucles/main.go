package main

import "fmt"

func main() {
	numeros := 100
	c := 0

	for numeros > 0 {
		numeros -= 1
		c++
	}

	fmt.Println(numeros)

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
