package main

import (
	"fmt"
	"strings"
)

func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {
	repeat3 := repeat(3)
	result := repeat3("Alfonso")
	fmt.Println(result)

	repeat5 := repeat(5)
	result2 := repeat5("Alfonsos")
	fmt.Println(result2)
}
