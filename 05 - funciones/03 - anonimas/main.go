package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hola desde funcion anonima")
	}()

	myFunc := func(text string) string {
		message := fmt.Sprintf("Hola desde funcion anonima 2 %s", text)

		return message
	}

	myStringResult := myFunc("pachimix")

	fmt.Println(myStringResult)
}
