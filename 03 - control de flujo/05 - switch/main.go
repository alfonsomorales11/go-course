package main

import "fmt"

func main() {
	var vocal string

	fmt.Println("Ingresa una vocal: ")
	fmt.Scanln(&vocal)

	switch {
	case vocal == "a" || vocal == "A":
		fmt.Println(vocal, "es vocal")

	case vocal == "e" || vocal == "E":
		fmt.Println(vocal, "es vocal")

	case vocal == "i" || vocal == "I":
		fmt.Println(vocal, "es vocal")

	case vocal == "o" || vocal == "O":
		fmt.Println(vocal, "es vocal")

	case vocal == "u" || vocal == "U":
		fmt.Println(vocal, "es vocal")

	default:
		fmt.Println("No es vocal")

	}

	switch vocal {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vocal")
	default:
		fmt.Println("No es vocal")
	}
}
