package main

import "fmt"

func main() {
	dias := make(map[int]string)

	dias[1] = "lunes"
	dias[2] = "martes"
	dias[3] = "miercoles"
	dias[4] = "jueves"
	dias[5] = "viernes"
	dias[6] = "sabado"
	dias[7] = "domingo"

	fmt.Println(dias)

	delete(dias, 1)

	estudiantes := make(map[string][]int)

	estudiantes["Alfonso"] = []int{10, 20, 30}
	estudiantes["Pachimix"] = []int{1, 2, 3}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["Alfonso"][1])
}
