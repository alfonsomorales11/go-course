package main

import "fmt"

func sumar(a int, b int) int {
	return a + b
}

func division(a int, b int) string {
	cociente := a / b
	residuo := a % b
	resultado := fmt.Sprintf("El cociente es %d y el residuo es %d", cociente, residuo)
	return resultado
}

func venta(a float32) string {
	precio_venta := a * 1.18
	igv := a * 0.18
	resultado := fmt.Sprintf("El precio de venta es %f y el IGV es %f", precio_venta, igv)
	return resultado
}

func main() {

	var num1 int
	var num2 int

	// Parte uno
	fmt.Print("Pon un numero: ")
	fmt.Scan(&num1)

	fmt.Print("Pon otro numero: ")
	fmt.Scan(&num2)

	var result int = sumar(num1, num2)

	fmt.Printf("El resultado de %d + %d es = %d", num1, num2, result)

	// Parte dos
	fmt.Print("\nPon un numero: ")
	fmt.Scan(&num1)

	fmt.Print("Pon otro numero: ")
	fmt.Scan(&num2)

	var resultado string = division(num1, num2)

	fmt.Println(resultado)

	// Parte tres
	var precio float32
	fmt.Print("\nPon el precio de venta ")
	fmt.Scan(&precio)

	resultado = venta(precio)
	fmt.Println(resultado)
}
