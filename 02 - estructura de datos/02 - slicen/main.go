package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3}

	numeros = append(numeros, 10)

	fmt.Println(numeros)

	subNumeros := numeros[1:]
	numeros[1] = 100
	fmt.Println(subNumeros)

	meses := []string{"enero", "febrero", "marzo"}
	meses = append(meses, "Abril")
	fmt.Printf("Len %v, Cap %v, %p\n", len(meses), cap(meses), meses)
}
