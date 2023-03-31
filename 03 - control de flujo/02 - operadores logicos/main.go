package main

import "fmt"

func main() {
	fmt.Println(true && false)

	fmt.Println(true || false)

	b := 2
	r := b == 2 && !(4 > b)

	fmt.Println(r)
}
