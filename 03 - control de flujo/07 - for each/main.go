package main

import "fmt"

func main() {
	nombres := [...]string{"alfonso", "morales", "pachi"}

	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	for index, element := range nombres {
		fmt.Println(index, element)
	}

}
